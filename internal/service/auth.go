package service

import (
	"context"

	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/repository"
)

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (srv *AuthService) Registration(ctx context.Context, user domain.User) error {
	return nil
}
