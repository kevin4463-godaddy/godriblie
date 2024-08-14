package godriblie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func NewLockClient(dynamoDb DynamoDbProvider, table string, opts ...DribbleClientOption) *DribbleClient {
	client := &DribbleClient{
		TableName:        table,
		OwnerName:        uuid.New().String(),
		PartitionKeyName: defaultPartitionKeyName,
		DynamoDB:         dynamoDb,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
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

	if opt.provisionedThroughput != nil {
		createTableInput.ProvisionedThroughput = opt.provisionedThroughput
	}

	return dc.DynamoDB.CreateTable(ctx, createTableInput)
}
