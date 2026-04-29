package repository

import (
	"context"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/infrastructure/postgres"
)

type UserRepo struct {
	pool infrastructure_postgres.Pool
}

func NewUserRepository(pool infrastructure_postgres.Pool) *UserRepo {
	return &UserRepo{pool}
}

func (r *UserRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO talan.users (first_name, last_name, middle_name, email, phone_number, password, created_at) 
				VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP) RETURNING id, first_name, last_name, middle_name, email, phone_number, created_at, updated_at`

	row := r.pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.MiddleName, user.Email, user.PhoneNumber, user.Password)

	var userDomain domain.User
	err := row.Scan(
		&userDomain.ID,
		&userDomain.FirstName,
		&userDomain.LastName,
		&userDomain.MiddleName,
		&userDomain.Email,
		&userDomain.PhoneNumber,
		&userDomain.CreatedAt,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return userDomain, nil
}

func (r *UserRepo) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, first_name, last_name, middle_name, 
       				email, phone_number, created_at, updated_at
					FROM talan.users ORDER BY id LIMIT $1 OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User

		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return users, nil
}
