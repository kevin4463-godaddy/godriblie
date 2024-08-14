package godriblie

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

func (dc *DribbleClient) AcquireLock(ctx context.Context, pKey string, opts ...AcquireLockOption) (*utils.LockDto, error) {
	lockReq := &acquireLockOptions{
		PartitionKey: pKey,
	}

	for _, opt := range opts {
		opt(lockReq)
	}

	return dc.acquireLock(ctx, lockReq)
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

	return dc.upsertNewLock(ctx, opt.AdditionalAttributes, opt.PartitionKey, opt.DeleteOnRelease, opt.Data, item)
}

func (dc *DribbleClient) upsertNewLock(ctx context.Context,
	additionalAttributes map[string]types.AttributeValue,
	key string,
	deleteLockOnRelease bool,
	newLockData []byte,
	item map[string]types.AttributeValue) (*utils.LockDto, error) {

	//cond := expression.AttributeExists(expression.Name(dc.PartitionKeyName))

	//putItemXpress, err := expression.NewBuilder().WithCondition(cond).Build()
	//if err != nil {
	//	return nil, err
	//}
	log.Printf("putting %v", item)
	req := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(dc.TableName),
	}

	// log something here
	return dc.putLockItem(ctx, key, req)
}

func (dc *DribbleClient) putLockItem(ctx context.Context,
	key string,
	putItemRequest *dynamodb.PutItemInput) (*utils.LockDto, error) {

	_, err := dc.DynamoDB.PutItem(ctx, putItemRequest)
	if err != nil {
		return nil, err
	}

	response := &utils.LockDto{
		PartitionKey: key,
	}

	return response, nil
}

func (dc *DribbleClient) ReleaseLock(ctx context.Context, lockName string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"key":   &types.AttributeValueMemberS{Value: lockName},
			"owner": &types.AttributeValueMemberS{Value: dc.OwnerName},
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

	if !lockItem.DeleteOnRelease {
		updateInput := &dynamodb.UpdateItemInput{
			TableName: aws.String(dc.TableName),
			Key: map[string]types.AttributeValue{
				"key":   &types.AttributeValueMemberS{Value: lockName},
				"owner": &types.AttributeValueMemberS{Value: dc.OwnerName},
			},
			UpdateExpression:    aws.String("SET #ts = :timestamp, isReleased = :isReleased"),
			ConditionExpression: aws.String("owner = :owner"),
			ExpressionAttributeNames: map[string]string{
				"#ts": "timestamp",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":timestamp":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().Unix())},
				":isReleased": &types.AttributeValueMemberBOOL{Value: true},
				":owner":      &types.AttributeValueMemberS{Value: dc.OwnerName},
			},
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
			"key":   &types.AttributeValueMemberS{Value: lockName},
			"owner": &types.AttributeValueMemberS{Value: dc.OwnerName},
		},
		ConditionExpression: aws.String("owner = :owner"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":owner": &types.AttributeValueMemberS{Value: dc.OwnerName},
		},
	}

	_, err := dc.DynamoDB.DeleteItem(ctx, deleteInput)
	if err != nil {
		return fmt.Errorf("failed to delete lock: %v", err)
	}

	return nil
}

func (dc *DribbleClient) CheckLock(ctx context.Context, lockName string) (bool, string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"key":   &types.AttributeValueMemberS{Value: lockName},
			"owner": &types.AttributeValueMemberS{Value: dc.OwnerName},
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
