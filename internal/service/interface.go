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
}
