package services

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/qchart-app/service-tv-udf/internal/domain"
// 	"github.com/qchart-app/service-tv-udf/internal/infrastructure/cache"
// )

// type UserService struct {
// 	userRepo domain.UserRepository
// 	cacheobj cache.CacheClient
// 	ctx      context.Context
// }

// func NewUserService(ctxobj context.Context, userRepo domain.UserRepository, cache cache.CacheClient) *UserService {
// 	return &UserService{
// 		userRepo: userRepo,
// 		cacheobj: cache,
// 		ctx:      ctxobj,
// 	}
// }

// func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
// 	// Try to get the user from the cache
// 	user, err := s.getUserFromCache(id)
// 	if err == nil {
// 		return user, nil
// 	}

// 	// If the user is not in the cache, get it from the repository
// 	user, err = s.userRepo.FindByID(s.ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Cache the user for future requests
// 	err = s.cacheobj.Set(getCacheKey(id), user, time.Minute*10)
// 	if err != nil {
// 		// Don't return an error if caching fails
// 		// This is because caching is an optimization, not a requirement
// 		// If caching fails, the application can continue to function without it
// 	}

// 	return user, nil
// }

// func (s *UserService) getUserFromCache(id uint) (*domain.User, error) {
// 	cacheKey := getCacheKey(id)

// 	val, err := s.cacheobj.GetOBJ(cacheKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, ok := val.(*domain.User)
// 	if !ok {
// 		return nil, domain.ErrInvalidCacheValue
// 	}

// 	return user, nil
// }

// func getCacheKey(id uint) string {
// 	return fmt.Sprintf("user:%d", id)
// }
