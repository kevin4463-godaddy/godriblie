package godriblie

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kevin4463-godaddy/godriblie/internal/utils"
)

func (dc *DribbleClient) AcquireLock(ctx context.Context, lockName string, deleteOnRelease bool, data interface{}) error {
	now := time.Now().Unix()
	ttl := 30 * 24 * 60 * 60

	dataStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	item := Lock{
		PartitionKey:    defaultPartitionKeyName,
		Owner:           dc.OwnerName,
		Timestamp:       now,
		ExpTime:         now + int64(ttl),
		DeleteOnRelease: deleteOnRelease,
		IsReleased:      false,
		Data:            string(dataStr),
	}

	av, err := utils.MarshalLockItem(item)

}
