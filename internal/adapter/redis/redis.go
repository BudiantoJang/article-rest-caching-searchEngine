package redis

import (
	"context"
	"jang-article/internal/port"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Client *redis.Client
}

func New(client *redis.Client) port.CacheRepository {
	return &Client{client}
}

func (rc *Client) SetIfNotExist(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	return rc.Client.SetNX(ctx, key, value, ttl).Result()
}

func (rc *Client) CheckHealth(ctx context.Context) (string, error) {
	return rc.Client.Ping(ctx).Result()
}
