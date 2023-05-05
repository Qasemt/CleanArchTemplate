package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/qchart-app/service-tv-udf/internal/domain"
	"github.com/qchart-app/service-tv-udf/internal/infrastructure/cache"
)

type UserService interface {
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(userID int) error
	GetUserByID(userID int) (*domain.User, error)
}

type userServiceRedis struct {
	cache_client cache.CacheClient
}

func NewUserServiceRedis(cache *cache.CacheClient) UserService {

	return &userServiceRedis{cache_client: *cache}
}

func (s *userServiceRedis) CreateUser(user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := getUserKey(user.ID)

	err = s.cache_client.Set(context.Background(), key, string(data), 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceRedis) UpdateUser(user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := getUserKey(user.ID)

	err = s.cache_client.Set(context.Background(), key, string(data), 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceRedis) DeleteUser(userID int) error {
	key := getUserKey(userID)

	err := s.cache_client.Del(context.Background(), key)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceRedis) GetUserByID(userID int) (*domain.User, error) {
	key := getUserKey(userID)

	data, err := s.cache_client.GetOBJ(context.Background(), key)
	if err == redis.Nil {
		return nil, domain.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return data.(*domain.User), nil
}

func getUserKey(userID int) string {
	return fmt.Sprintf("user:%d", userID)
}
