package domain

import "fmt"

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
