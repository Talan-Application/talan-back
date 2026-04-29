package repository

import (
	"context"

	"github.com/Talan-Application/talan-back/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error)
}
