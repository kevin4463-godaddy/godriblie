package godriblie

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func WithPartitionKeyName(pkn string) DribbleClientOption {
	return func(dc *DribbleClient) { dc.PartitionKeyName = pkn }
}

func WithOwnerName(name string) DribbleClientOption {
	return func(dc *DribbleClient) { dc.OwnerName = name }
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

func WithData(b []byte) AcquireLockOption {
	return func(opt *acquireLockOptions) {
		opt.Data = b
	}
}

// ReplaceData will force the new content to be stored in the key.
func ReplaceData() AcquireLockOption {
	return func(opt *acquireLockOptions) {
		opt.ReplaceData = true
	}
}

// FailIfLocked will not retry to acquire the lock, instead returning.
func FailIfLocked() AcquireLockOption {
	return func(opt *acquireLockOptions) {
		opt.FailIfLocked = true
	}
}

// WithDeleteLockOnRelease defines whether or not the lock should be deleted
// when Close() is called on the resulting LockItem will force the new content
// to be stored in the key.
func WithDeleteLockOnRelease() AcquireLockOption {
	return func(opt *acquireLockOptions) {
		opt.DeleteOnRelease = true
	}
}

// WithAdditionalAttributes stores some additional attributes with each lock.
// This can be used to add any arbitrary parameters to each lock row.
func WithAdditionalAttributes(attr map[string]types.AttributeValue) AcquireLockOption {
	return func(opt *acquireLockOptions) {
		opt.AdditionalAttributes = attr
	}
}
