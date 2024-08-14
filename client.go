package godriblie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewLockClient(dynamoDb DynamoDbProvider, table string, opts ...DribbleClientOption) *DribbleClient {
	client := &DribbleClient{
		TableName:        table,
		OwnerName:        "default:owneroflock",
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

func WithCustomPartitionKeyName(s string) CreateTableOption {
	return func(opt *createDynamoDBTableOptions) {
		opt.partitionKeyName = s
	}
}

func WithProvisionedThroughput(provisionedThroughput *types.ProvisionedThroughput) CreateTableOption {
	return func(opt *createDynamoDBTableOptions) {
		opt.billingMode = types.BillingModeProvisioned
		opt.provisionedThroughput = provisionedThroughput
	}
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
