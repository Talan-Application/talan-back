package service

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (s *AuthService) Registration(ctx context.Context, userReq domain.User) error {
	_, err := s.userRepository.GetUserByEmail(ctx, userReq.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", userReq.Email)
	}
	if !errors.Is(err, core_errors.ErrNotFound) {
		return fmt.Errorf("lookup error: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 12)
	if err != nil {
		return fmt.Errorf("bcrypt error: %w", err)
	}
	userReq.Password = string(hash)

	if _, err := s.userRepository.CreateUser(ctx, userReq); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
