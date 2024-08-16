package utils

type LockDto struct {
	PartitionKey    string `dynamodbav:"key"`
	Owner           string `dynamodbav:"owner"`
	Timestamp       int64  `dynamodbav:"timestamp"`
	ExpTime         int64  `dynamodbav:"expTime"`
	DeleteOnRelease bool   `dynamodbav:"deleteOnRelease"`
	IsReleased      bool   `dynamodbav:"isReleased"`
	Data            []byte `dynamodbav:"data"`
}
