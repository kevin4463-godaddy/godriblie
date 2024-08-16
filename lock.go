package godriblie

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

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

func (dc *DribbleClient) validateLock(ctx context.Context, pKey string) error {
	query := &dynamodb.QueryInput{
		TableName:              aws.String(dc.TableName),
		KeyConditionExpression: aws.String("#key = :pKey"),
		ExpressionAttributeNames: map[string]string{
			"#key": "key",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pKey": &types.AttributeValueMemberS{Value: pKey},
		},
	}

	q, err := dc.DynamoDB.Query(ctx, query)
	if err != nil {
		log.Printf("failed to query: %v", err)
		return err
	}

	var locks []utils.LockDto
	locks, err = utils.UnmarshalLockItemList(q.Items)
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
