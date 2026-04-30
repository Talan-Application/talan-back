package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/Talan-Application/talan-back/internal/infrastructure/postgres"
	"github.com/jackc/pgx/v5"
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

	query := `INSERT INTO talan.users (first_name, last_name, middle_name,
                         email, phone_number, password, role, created_at) 
             VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP) 
             RETURNING id, first_name, last_name, middle_name, email, phone_number, role, created_at, updated_at`

	row := r.pool.QueryRow(ctx, query,
		user.FirstName, user.LastName, user.MiddleName,
		user.Email, user.PhoneNumber, user.Password, user.Role)
	var userDomain domain.User
	err := row.Scan(
		&userDomain.ID,
		&userDomain.FirstName,
		&userDomain.LastName,
		&userDomain.MiddleName,
		&userDomain.Email,
		&userDomain.PhoneNumber,
		&userDomain.Role,
		&userDomain.CreatedAt,
		&userDomain.UpdatedAt,
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
       				email, phone_number, role, created_at, updated_at
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
			&user.Role,
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

func (r *UserRepo) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, first_name, last_name, middle_name,
       			email, phone_number, role, created_at, updated_at
             FROM talan.users WHERE id = $1`

	row := r.pool.QueryRow(ctx, query, id)
	var user domain.User
	err := row.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.MiddleName,
		&user.Email, &user.PhoneNumber, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}
	return user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id, first_name, last_name, middle_name, 
			email, phone_number, role, created_at, updated_at
				FROM talan.users WHERE email = $1`

	row := r.pool.QueryRow(ctx, query, email)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with email=%s: %w",
				email,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}

	return user, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM talan.users WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id=%d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE talan.users SET first_name=$1, last_name=$2,
                       middle_name=$3, phone_number=$4, updated_at=CURRENT_TIMESTAMP
                   WHERE id = $5 RETURNING id, first_name, last_name, middle_name, email,
                       phone_number, role, created_at, updated_at`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.PhoneNumber,
		id,
	)

	var updatedUser domain.User
	err := row.Scan(
		&updatedUser.ID,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.MiddleName,
		&updatedUser.Email,
		&updatedUser.PhoneNumber,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id=%d: %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}

	return updatedUser, nil
}
