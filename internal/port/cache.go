package port

import (
	"context"
	"time"
)

type CacheRepository interface {
	SetIfNotExist(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error)
	CheckHealth(ctx context.Context) (string, error)
}
