package request

import (
	"errors"

	"github.com/Talan-Application/talan-back/internal/domain"
)

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}

func DomainFromCreateUserDto(dto CreateUserRequest) domain.User {
	userRole, err := domain.ParseRole(dto.Role)
	if err != nil {
		userRole = domain.RoleStudent
	}

	return domain.User{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		MiddleName: &dto.MiddleName,
		Password:   dto.Password,
		Email:      dto.Email,
		Role:       userRole,
	}
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	PhoneNumber string `json:"phone_number"`
}

func (s UpdateUserRequest) ValidateUpdateUserRequest() error {
	if s.FirstName == "" {
		return errors.New("first name can not be empty")
	} else if s.LastName == "" {
		return errors.New("last name can not be empty")
	}

	return nil
}

func DomainFromUpdateUserDto(dto UpdateUserRequest) domain.User {
	return domain.User{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		MiddleName:  &dto.MiddleName,
		PhoneNumber: &dto.PhoneNumber,
	}
}
