package utils

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MarshalLockItem(item LockDto) (map[string]types.AttributeValue, error) {
	av := map[string]types.AttributeValue{
		"key":             &types.AttributeValueMemberS{Value: item.PartitionKey},
		"owner":           &types.AttributeValueMemberS{Value: item.Owner},
		"timestamp":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.Timestamp)},
		"expTime":         &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.ExpTime)},
		"deleteOnRelease": &types.AttributeValueMemberBOOL{Value: item.DeleteOnRelease},
		"isReleased":      &types.AttributeValueMemberBOOL{Value: item.IsReleased},
		"data":            &types.AttributeValueMemberS{Value: string(item.Data)},
	}

	return av, nil
}

func UnmarshalLockItem(item map[string]types.AttributeValue) (LockDto, error) {
	timestamp, err := strconv.ParseInt(item["timestamp"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return LockDto{}, err
	}
	expTime, err := strconv.ParseInt(item["expTime"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return LockDto{}, err
	}
	deleteOnRelease := item["deleteOnRelease"].(*types.AttributeValueMemberBOOL).Value
	isReleased := item["isReleased"].(*types.AttributeValueMemberBOOL).Value
	data := item["data"].(*types.AttributeValueMemberS).Value
	partitionKey := item["key"].(*types.AttributeValueMemberS).Value
	owner := item["owner"].(*types.AttributeValueMemberS).Value

	result := LockDto{
		PartitionKey:    partitionKey,
		Owner:           owner,
		Timestamp:       timestamp,
		ExpTime:         expTime,
		DeleteOnRelease: deleteOnRelease,
		IsReleased:      isReleased,
		Data:            []byte(data),
	}

	return result, nil
}
