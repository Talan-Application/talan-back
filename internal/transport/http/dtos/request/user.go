package request

import "github.com/Talan-Application/talan-back/internal/domain"

type CreateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func DomainFromCreateUserDto(dto CreateUserRequest) domain.User {
	return domain.User{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		MiddleName:  &dto.MiddleName,
		Password:    dto.Password,
		Email:       dto.Email,
		PhoneNumber: &dto.PhoneNumber,
	}
}
