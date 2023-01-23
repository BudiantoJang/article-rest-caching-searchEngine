package port

import (
	"context"
	"time"
)

type CacheRepository interface {
	SetIfNotExist(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error)
	Get(ctx context.Context, key string) (string, error)
	CheckHealth(ctx context.Context) (string, error)
}
