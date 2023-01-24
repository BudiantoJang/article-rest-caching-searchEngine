package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	Client *redis.Client
}

func New(client *redis.Client) *Client {
	return &Client{client}
}

func (rc *Client) SetIfNotExist(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	json, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	return rc.Client.SetNX(ctx, key, json, ttl).Result()
}

func (rc *Client) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (rc *Client) CheckHealth(ctx context.Context) (string, error) {
	return rc.Client.Ping(ctx).Result()
}
