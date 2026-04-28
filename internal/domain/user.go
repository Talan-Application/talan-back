package domain

import "time"

type User struct {
	ID          int        `db:"id"`
	FirstName   string     `db:"first_name"`
	LastName    string     `db:"last_name"`
	MiddleName  *string    `db:"middle_name"`
	PhoneNumber *string    `db:"phone_number"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	IsVerified  bool       `db:"is_verified"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
