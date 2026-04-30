package service

import (
	"context"

	"github.com/Talan-Application/talan-back/internal/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error)
}

type IAuthService interface {
	Registration(ctx context.Context, user domain.User) error
	Authenticate(ctx context.Context, email string, password string) (domain.Token, error)
}
