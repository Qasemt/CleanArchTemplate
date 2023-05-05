package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/qchart-app/service-tv-udf/internal/domain"
)

type DriverRedis struct {
	client *redis.Client
}

func NewRedisClient(addr, password string, db int) (CacheClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Check the connection to the Redis server
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return &DriverRedis{client}, nil
}

func (r *DriverRedis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (c *DriverRedis) GetOBJ(ctx context.Context, key string) (interface{}, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, domain.ErrCacheMiss
	} else if err != nil {
		return nil, err
	}

	return []byte(value), nil
}
func (r *DriverRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *DriverRedis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
func (c *DriverRedis) NewSubscriber() Subscriber {
	return &RedisSubscriber{
		Client: c.client,
	}
}
