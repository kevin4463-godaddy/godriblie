package godriblie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

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

	return dc.DynamoDB.CreateTable(ctx, createTableInput)
}
