package cache

import (
	"context"
	"time"
)

type CacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	GetOBJ(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	NewSubscriber() Subscriber
}
