package model

import "context"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id uint) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}
type UserUseCase interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uint) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, uid uint) error
}
