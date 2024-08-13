package godriblie

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

func (dc *DribbleClient) AcquireLock(ctx context.Context, lockName string, deleteOnRelease bool, data interface{}) error {
	now := time.Now().Unix()
	ttl := 30 * 24 * 60 * 60

	dataStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	item := Lock{
		PartitionKey:    lockName,
		Owner:           dc.OwnerName,
		Timestamp:       now,
		ExpTime:         now + int64(ttl),
		DeleteOnRelease: deleteOnRelease,
		IsReleased:      false,
		Data:            string(dataStr),
	}

	av, err := utils.MarshalLockItem(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(dc.TableName),
		ConditionExpression: aws.String("attribute_not_exists(PartitionKey)"),
	}

	_, err = dc.DynamoDB.PutItem(ctx, input)
	if err != nil {
		var cfe *types.ConditionalCheckFailedException
		if errors.As(err, &cfe) {
			return errors.New("lock is already help by another owner")
		}
		return fmt.Errorf("fialed to acquire lock: %v", err)
	}

	return nil
}

func (dc *DribbleClient) ReleaseLock(ctx context.Context, lockName string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dc.TableName),
		Key: map[string]types.AttributeValue{
			"LockName": &types.AttributeValueMemberS{Value: lockName},
			"Owner":    &types.AttributeValueMemberS{Value: dc.OwnerName},
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
				"LockName": &types.AttributeValueMemberS{Value: lockName},
				"Owner":    &types.AttributeValueMemberS{Value: dc.OwnerName},
			},
			UpdateExpression:    aws.String("SET #ts = :timestamp, IsReleased = :isReleased"),
			ConditionExpression: aws.String("Owner = :owner"),
			ExpressionAttributeNames: map[string]string{
				"#ts": "Timestamp",
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
			"LockName": &types.AttributeValueMemberS{Value: lockName},
			"Owner":    &types.AttributeValueMemberS{Value: dc.OwnerName},
		},
		ConditionExpression: aws.String("Owner = :owner"),
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
			"LockName": &types.AttributeValueMemberS{Value: lockName},
			"Owner":    &types.AttributeValueMemberS{Value: dc.OwnerName},
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

	return lockItem.IsReleased, lockItem.Data, nil
}
