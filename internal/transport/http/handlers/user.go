package handlers

import (
	"net/http"
	"strconv"

	core_errors "github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/service"
	"github.com/Talan-Application/talan-back/internal/transport/http/dtos/request"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core_errors.HandleError(c, err)
		return
	}

	res, err := h.userService.CreateUser(c.Request.Context(), request.DomainFromCreateUserDto(req))
	if err != nil {
		core_errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		core_errors.HandleError(c, err)
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		core_errors.HandleError(c, err)
		return
	}

	res, err := h.userService.GetUsers(c.Request.Context(), &limit, &offset)
	if err != nil {
		core_errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
