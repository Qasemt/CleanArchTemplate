package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Subscriber interface {
	Subscribe(ctx context.Context, channel string) (<-chan string, error)
}

type RedisSubscriber struct {
	Client *redis.Client
}

func (s *RedisSubscriber) Subscribe(ctx context.Context, channel string) (<-chan string, error) {

	pubsub := s.Client.PSubscribe(ctx, channel)
	ch := make(chan string)
	go func() {
		defer func() {
			pubsub.Close()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := pubsub.ReceiveMessage(ctx)
				if err != nil {
					continue
				}
				ch <- msg.Payload
			}
		}
	}()

	return ch, nil
}
