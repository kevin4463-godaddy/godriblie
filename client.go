package godriblie

import (
	"context"
	"errors"
	"fmt"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func NewLockClient(dynamoDb DynamoDbProvider, table string, opts ...DribbleClientOption) *DribbleClient {
	client := &DribbleClient{
		TableName:        table,
		OwnerName:        uuid.New().String(),
		PartitionKeyName: defaultPartitionKeyName,
		DynamoDB:         dynamoDb,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (dc *DribbleClient) CreateTable(ctx context.Context, tableName string, opts ...CreateTableOption) (*dynamodb.CreateTableOutput, error) {
	createTableOptions := &createDynamoDBTableOptions{
		tableName:        tableName,
		billingMode:      "PAY_PER_REQUEST",
		partitionKeyName: dc.PartitionKeyName,
	}
	for _, opt := range opts {
		opt(createTableOptions)
	}

	return dc.createTable(ctx, createTableOptions)
}

func (dc *DribbleClient) AcquireLock(ctx context.Context, pKey string, opts ...AcquireLockOption) (*utils.LockDto, error) {

	err := dc.validateLock(ctx, pKey)
	if err != nil {
		return nil, err
	}

	lockOptions := &acquireLockOptions{
		PartitionKey: pKey,
	}

	for _, opt := range opts {
		opt(lockOptions)
	}

	return dc.acquireLock(ctx, lockOptions)
}

func (dc *DribbleClient) ReleaseLock(ctx context.Context, lockName string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: lockName},
		},
	}

	result, err := dc.DynamoDB.GetItem(ctx, input)
	if err != nil {
		return fmt.Errorf("fialed to get lock item: %v", err)
	}

	if result.Item == nil {
		return errors.New("lock item not found")
	}

	lockItem, err := utils.UnmarshalLockItem(result.Item)
	if err != nil {
		return err
	}

	// only release/delete if client is current owner
	if !lockItem.DeleteOnRelease {
		updateInput := &dynamodb.UpdateItemInput{
			TableName: aws.String(dc.TableName),
			Key: map[string]types.AttributeValue{
				"key": &types.AttributeValueMemberS{Value: lockName},
			},
			UpdateExpression: aws.String("SET #ts = :timestamp, isReleased = :isReleased"),
			ExpressionAttributeNames: map[string]string{
				"#ts":    "timestamp",
				"#owner": "owner",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":owner":      &types.AttributeValueMemberS{Value: dc.OwnerName},
				":timestamp":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().Unix())},
				":isReleased": &types.AttributeValueMemberBOOL{Value: true},
			},
			ConditionExpression: aws.String("#owner = :owner"),
		}
		_, err = dc.DynamoDB.UpdateItem(ctx, updateInput)
		if err != nil {
			return fmt.Errorf("fialed to update lock release: %v", err)
		}

		return nil
	}

	return dc.DeleteOnReleaseLock(ctx, lockName)
}

func (dc *DribbleClient) DeleteOnReleaseLock(ctx context.Context, lockName string) error {
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: lockName},
		},
		ExpressionAttributeNames: map[string]string{
			"#owner": "owner",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":owner":  &types.AttributeValueMemberS{Value: dc.OwnerName},
			":delete": &types.AttributeValueMemberBOOL{Value: true},
		},
		ConditionExpression: aws.String("deleteOnRelease = :delete and #owner = :owner"),
	}

	_, err := dc.DynamoDB.DeleteItem(ctx, deleteInput)
	if err != nil {
		return fmt.Errorf("failed to delete lock: %v", err)
	}

	return nil
}

func (dc *DribbleClient) CheckLock(ctx context.Context, lockName string) (bool, string, error) {
	input := &dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: lockName},
		},
	}

	result, err := dc.DynamoDB.GetItem(ctx, input)
	if err != nil {
		return false, "", fmt.Errorf("fialed to check lock: %v", err)
	}

	if result.Item == nil {
		return false, "", nil
	}

	lockItem, err := utils.UnmarshalLockItem(result.Item)
	if err != nil {
		return false, "", err
	}

	return lockItem.IsReleased, string(lockItem.Data), nil
}

func (dc *DribbleClient) createTable(ctx context.Context, opt *createDynamoDBTableOptions) (*dynamodb.CreateTableOutput, error) {
	keySchema := []types.KeySchemaElement{
		{
			AttributeName: aws.String(opt.partitionKeyName),
			KeyType:       types.KeyTypeHash,
		},
	}

	attributeDefinitions := []types.AttributeDefinition{
		{
			AttributeName: aws.String(opt.partitionKeyName),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}

	createTableInput := &dynamodb.CreateTableInput{
		TableName:            aws.String(dc.TableName),
		KeySchema:            keySchema,
		BillingMode:          opt.billingMode,
		AttributeDefinitions: attributeDefinitions,
	}

	if opt.provisionedThroughput != nil {
		createTableInput.ProvisionedThroughput = opt.provisionedThroughput
	}

	return dc.DynamoDB.CreateTable(ctx, createTableInput)
}
