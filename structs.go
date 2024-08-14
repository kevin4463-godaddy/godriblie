package godriblie

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

const (
	defaultPartitionKeyName = "key"
)

type DribbleClient struct {
	TableName        string
	OwnerName        string
	PartitionKeyName string
	DynamoDB         DynamoDbProvider
}

type DribbleClientOption func(*DribbleClient)

type createDynamoDBTableOptions struct {
	billingMode           types.BillingMode
	provisionedThroughput *types.ProvisionedThroughput
	tableName             string
	partitionKeyName      string
	sortKeyName           string
	tags                  []types.Tag
}

type CreateTableOption func(*createDynamoDBTableOptions)
