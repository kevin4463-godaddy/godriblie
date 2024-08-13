package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie"
)

func MarshalLockItem(item godriblie.Lock) (map[string]types.AttributeValue, error) {
	data, err := json.Marshal(item.Data)
	if err != nil {
		return nil, err
	}

	av := map[string]types.AttributeValue{
		"LockName":        &types.AttributeValueMemberS{Value: item.PartitionKey},
		"Owner":           &types.AttributeValueMemberS{Value: item.Owner},
		"Timestamp":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.Timestamp)},
		"ExpTime":         &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", item.ExpTime)},
		"DeleteOnRelease": &types.AttributeValueMemberBOOL{Value: item.DeleteOnRelease},
		"IsReleased":      &types.AttributeValueMemberBOOL{Value: item.IsReleased},
		"Data":            &types.AttributeValueMemberS{Value: string(data)},
	}

	return av, nil
}

func UnmarshalLockItem(item map[string]types.AttributeValue) (godriblie.Lock, error) {
	timestamp, err := strconv.ParseInt(item["Timestamp"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return godriblie.Lock{}, err
	}
	expTime, err := strconv.ParseInt(item["ExpTime"].(*types.AttributeValueMemberN).Value, 10, 64)
	if err != nil {
		return godriblie.Lock{}, err
	}
	deleteOnRelease := item["DeleteOnRelease"].(*types.AttributeValueMemberBOOL).Value
	isReleased := item["IsReleased"].(*types.AttributeValueMemberBOOL).Value
	data := item["Data"].(*types.AttributeValueMemberS).Value
	partitionKey := item["LockName"].(*types.AttributeValueMemberS).Value
	owner := item["Owner"].(*types.AttributeValueMemberS).Value

	result := godriblie.Lock{
		PartitionKey:    partitionKey,
		Owner:           owner,
		Timestamp:       timestamp,
		ExpTime:         expTime,
		DeleteOnRelease: deleteOnRelease,
		IsReleased:      isReleased,
		Data:            data,
	}

	return result, nil
}
