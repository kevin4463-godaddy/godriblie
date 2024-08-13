package godriblie

import "github.com/kevin4463-godaddy/godriblie/internal/utils"

func NewLockClient(dynamoDb DynamoDbProvider, table string, opts ...DribbleClientOption) *DribbleClient {
	client := &DribbleClient{
		TableName:        table,
		OwnerName:        utils.GenerateRandString(),
		PartitionKeyName: defaultPartitionKeyName,
		DynamoDB:         dynamoDb,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithPartitionKeyName(pkn string) DribbleClientOption {
	return func(dc *DribbleClient) { dc.PartitionKeyName = pkn }
}

func WithOwnerName(name string) DribbleClientOption {
	return func(dc *DribbleClient) { dc.OwnerName = name }
}
