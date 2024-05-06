package rpc

import "time"

const (
	NatsDefaultBucketName      = "minotaur-rpc-services"
	NatsDefaultBucketDesc      = "Minotaur RPC Services Management Bucket, used for storing service information."
	NatsDefaultBkvBucketPrefix = "minotaur-rpc-services"
	NatsDefaultTTL             = 10 * time.Second
	NatsDefaultKeepAlive       = 3 * time.Second
)
