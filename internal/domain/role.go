package domain

import (
	"database/sql/driver"
	"fmt"
)

type UserRole int

const (
	RoleStudent UserRole = iota // 0
	RoleCurator                 // 1
	RoleTeacher                 // 2
	RoleAdmin                   // 3
)

func (r UserRole) String() string {
	return [...]string{"student", "curator", "teacher", "admin"}[r]
}

func (u *User) HasRequiredRole(targetRole UserRole) bool {
	return u.Role == targetRole
}

func ParseRole(roleStr string) (UserRole, error) {
	switch roleStr {
	case "student":
		return RoleStudent, nil
	case "curator":
		return RoleCurator, nil
	case "teacher":
		return RoleTeacher, nil
	case "admin":
		return RoleAdmin, nil
	default:
		return RoleStudent, fmt.Errorf("invalid role: %s", roleStr)
	}
}

func (r UserRole) Value() (driver.Value, error) {
	return r.String(), nil
}

func (r *UserRole) Scan(value interface{}) error {
	if value == nil {
		*r = RoleStudent
		return nil
	}

	sv, ok := value.(string)
	if !ok {
		// Some drivers return []byte, so we handle that too
		bv, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("cannot scan %T into UserRole", value)
		}
		sv = string(bv)
	}

	role, err := ParseRole(sv)
	if err != nil {
		return err
	}
	*r = role
	return nil
}
