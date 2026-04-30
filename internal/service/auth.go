package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	core_errors "github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/repository"
	"github.com/golang-jwt/jwt/v5"
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

func (s *AuthService) Authenticate(ctx context.Context, email string, password string) (domain.Token, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.Token{}, fmt.Errorf("get user: %w", err)
	}
	if !user.IsVerified {
		return domain.Token{}, fmt.Errorf(
			"user with email=%s not verified: %w",
			email,
			core_errors.ErrUnauthorized,
		)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.Token{}, fmt.Errorf("wrong password: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": int(user.Role),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	})

	secret := []byte("mysecretkeyhere")
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return domain.Token{}, fmt.Errorf("create token: %w", err)
	}

	return domain.Token{AccessToken: tokenStr}, nil
}
