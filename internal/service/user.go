package service

import (
	"context"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/repository"
)

type UserSrv struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserSrv {
	return &UserSrv{userRepository}
}

func (s *UserSrv) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.userRepository.CreateUser(ctx, user)
}

func (s *UserSrv) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must not be negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must not be negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf(
			"get users from repository: %w",
			err,
		)
	} else if len(users) == 0 {
		return nil, fmt.Errorf(
			"users: %w", core_errors.ErrNotFound)
	}

	return users, nil
}

func (s *UserSrv) GetUserById(ctx context.Context, id int) (domain.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}
	return user, nil
}

func (s *UserSrv) DeleteUser(ctx context.Context, id int) error {
	if err := s.userRepository.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
