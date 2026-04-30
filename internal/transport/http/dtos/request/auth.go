package request

import (
	"fmt"

	core_errors "github.com/Talan-Application/talan-back/internal/core/errors"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *LoginRequest) Validate() error {
	if s.Email == "" {
		return fmt.Errorf(
			"email is required: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if s.Password == "" {
		return fmt.Errorf(
			"password is required: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
