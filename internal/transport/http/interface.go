package http

import "github.com/Talan-Application/talan-back/internal/service"

type UserService interface {
	service.UserService
}

type AuthService interface {
	service.IAuthService
}
