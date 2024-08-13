package godriblie

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

type Lock struct {
	PartitionKey    string
	Owner           string
	Timestamp       int64
	ExpTime         int64
	DeleteOnRelease bool
	IsReleased      bool
	Data            string
	//AdditionalAttributes map[string]types.AttributeValue
}
