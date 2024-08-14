package utils

type LockDto struct {
	PartitionKey    string
	Owner           string
	Timestamp       int64
	ExpTime         int64
	DeleteOnRelease bool
	IsReleased      bool
	Data            []byte
}
