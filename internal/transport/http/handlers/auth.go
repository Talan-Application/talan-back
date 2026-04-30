package handlers

import (
	"fmt"
	"net/http"

	"github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/service"
	"github.com/Talan-Application/talan-back/internal/transport/http/dtos/request"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Registration(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core_errors.HandleError(c, fmt.Errorf(
			"bind json: %w",
			core_errors.ErrInvalidArgument,
		))
		return
	}

	if err := h.authService.Registration(
		c.Request.Context(),
		request.DomainFromCreateUserDto(req),
	); err != nil {
		core_errors.HandleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}
