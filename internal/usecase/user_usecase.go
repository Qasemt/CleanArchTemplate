package usecase

import (
	"context"

	domain "github.com/qchart-app/service-tv-udf/internal/domain"
)

type userUseCase struct {
	userRepo    domain.UserRepository
	userService UserService
}

func NewUserUseCase(userRepo domain.UserRepository, srv UserService) domain.UserUseCase {

	return &userUseCase{userRepo: userRepo, userService: srv}
}

func (u *userUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	// Perform input validation
	if user.FirstName == "" || user.Email == "" {
		return domain.ErrBadParamInput
	}

	// Create user in the repository
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := u.userService.GetUserByID(int(id))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	// Perform input validation
	if user.ID == 0 || user.FirstName == "" || user.Email == "" {
		return domain.ErrBadParamInput
	}

	// Update user in the repository
	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	user, err := u.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	err = u.userRepo.Delete(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
