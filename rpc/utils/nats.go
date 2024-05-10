package utils

import (
	"errors"
	"github.com/nats-io/nats.go"
	"time"
)

func InitNatsBucket(js nats.JetStreamContext, bucketName, bucketDesc string, ttl time.Duration) (nats.KeyValue, error) {
	kv, err := js.KeyValue(bucketName)
	if err != nil {
		if !errors.Is(err, nats.ErrBucketNotFound) {
			return nil, err
		}
	}

	if kv == nil {
		kv, err = js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:      bucketName,
			Description: bucketDesc,
			History:     9,
			TTL:         ttl,
			Storage:     nats.FileStorage,
		})
		if err != nil {
			return nil, err
		}
	}
	return kv, nil
}
