package repository

import (
	"context"

	"github.com/qchart-app/service-tv-udf/internal/domain"
	"github.com/qchart-app/service-tv-udf/internal/infrastructure/database"
	"gorm.io/gorm"
)

type userRepository struct {
	db *database.GormDB
}

func NewGormUserRepository(db *database.GormDB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.GORM.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
func (r *userRepository) Delete(ctx context.Context, user *domain.User) error {
	return r.db.GORM.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.GORM.WithContext(ctx).Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (r *userRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.GORM.WithContext(ctx).First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return r.db.GORM.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Save(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
