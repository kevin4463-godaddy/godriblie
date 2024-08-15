package godriblie

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

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

func (dc *DribbleClient) acquireLock(ctx context.Context, opt *acquireLockOptions) (*utils.LockDto, error) {
	lockOptions := utils.LockDto{
		PartitionKey:    opt.PartitionKey,
		Owner:           dc.OwnerName,
		Timestamp:       time.Now().Unix(),
		ExpTime:         30 * 24 * 60 * 60,
		DeleteOnRelease: opt.DeleteOnRelease,
		IsReleased:      false,
		Data:            opt.Data,
	}

	item, err := utils.MarshalLockItem(lockOptions)
	if err != nil {
		return nil, err
	}

	req := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(dc.TableName),
	}

	_, err = dc.DynamoDB.PutItem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &lockOptions, nil
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

func (dc *DribbleClient) validateLock(ctx context.Context, pKey string) error {
	query := &dynamodb.QueryInput{
		TableName:              aws.String(dc.TableName),
		KeyConditionExpression: aws.String("key = :pKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pKey": &types.AttributeValueMemberS{Value: pKey},
		},
	}

	q, err := dc.DynamoDB.Query(ctx, query)
	if err != nil {
		log.Printf("failed to query: %v", err)
		return err
	}

	var locks []*utils.LockDto
	err = attributevalue.UnmarshalListOfMaps(q.Items, &locks)
	if err != nil {
		log.Printf("failed to unmarshal locks: %v", err)
		return err
	}

	for _, lock := range locks {
		if lock.Owner != dc.OwnerName && !lock.IsReleased {
			return errors.New("lock is in use by another and not released")
		}
	}

	return nil
}
