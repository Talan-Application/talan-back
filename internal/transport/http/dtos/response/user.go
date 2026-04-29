package response

import (
	"time"

	"github.com/Talan-Application/talan-back/internal/domain"
)

type UserResponse struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	MiddleName  *string    `json:"middle_name"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	Role        string     `json:"role"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func UserResponseFromDomain(domain domain.User) UserResponse {
	return UserResponse{
		ID:          domain.ID,
		FirstName:   domain.FirstName,
		LastName:    domain.LastName,
		MiddleName:  domain.MiddleName,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		Role:        domain.Role.String(),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
