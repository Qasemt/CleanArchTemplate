package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/qchart-app/service-tv-udf/internal/infrastructure/cache"
)

type RedisMSGService struct {
	subscriber cache.Subscriber
}

type MessageService interface {
	ListenForMessages(ctx context.Context, channel string) error
}

func NewRedisMSGService(subscriber cache.Subscriber) MessageService {
	return &RedisMSGService{subscriber: subscriber}
}

func (s *RedisMSGService) ListenForMessages(ctx context.Context, channel string) error {

	msgCh, err := s.subscriber.Subscribe(ctx, channel)
	if err != nil {
		fmt.Printf("srv exit")
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgCh:
				log.Println("Received message from redis:", msg)

				// Parse the incoming JSON message
				var jsonData map[string]interface{}
				err := json.Unmarshal([]byte(msg), &jsonData)
				if err != nil {
					log.Println("Error parsing JSON message:", err)
					continue
				}

				// Call the message handler function with the parsed JSON data
			}
		}
	}()

	return nil
}
