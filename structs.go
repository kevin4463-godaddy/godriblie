package godriblie

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	attrData                = "data"
	attrOwnerName           = "owner"
	attrIsReleased          = "isReleased"
	defaultPartitionKeyName = "key"
)

var (
	dataAttr            = expression.Name(attrData)
	ownerNameAttr       = expression.Name(attrOwnerName)
	isReleasedAttr      = expression.Name(attrIsReleased)
	isReleasedAttrValue = expression.Value("1")
)

type DribbleClient struct {
	TableName        string
	OwnerName        string
	PartitionKeyName string
	DynamoDB         DynamoDbProvider
}

type createDynamoDBTableOptions struct {
	billingMode           types.BillingMode
	provisionedThroughput *types.ProvisionedThroughput
	tableName             string
	partitionKeyName      string
	sortKeyName           string
	tags                  []types.Tag
}

type acquireLockOptions struct {
	PartitionKey         string
	Data                 []byte
	ReplaceData          bool
	DeleteOnRelease      bool
	FailIfLocked         bool
	AdditionalAttributes map[string]types.AttributeValue
}

type getLockOptions struct {
	partitionKeyName     string
	deleteLockOnRelease  bool
	replaceData          bool
	data                 []byte
	additionalAttributes map[string]types.AttributeValue
	failIfLocked         bool
}

type DribbleClientOption func(*DribbleClient)
type CreateTableOption func(*createDynamoDBTableOptions)
type AcquireLockOption func(*acquireLockOptions)
