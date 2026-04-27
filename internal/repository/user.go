package repository

import (
	"context"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/domain"
)

type UserRepository struct {
	pool Pool
}

func NewUserRepository(pool Pool) *UserRepository {
	return &UserRepository{pool}
}

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO talan.users (first_name, last_name, middle_name, email, password, created_at) 
				VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP) RETURNING id, first_name, last_name, middle_name, email, created_at`

	row := r.pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.MiddleName, user.Email, user.Password)

	var userDomain domain.User
	err := row.Scan(
		&userDomain.ID,
		&userDomain.FirstName,
		&userDomain.LastName,
		&userDomain.MiddleName,
		&userDomain.Email,
		&userDomain.CreatedAt,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return userDomain, nil
}
