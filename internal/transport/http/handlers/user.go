package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/Talan-Application/talan-back/internal/service"
	"github.com/Talan-Application/talan-back/internal/transport/http/dtos/request"
	"github.com/Talan-Application/talan-back/internal/transport/http/dtos/response"
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

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core_errors.HandleError(c, fmt.Errorf(
			"convert id to int error: %w",
			core_errors.ErrInvalidArgument,
		))
		return
	}

	user, err := h.userService.GetUserById(c.Request.Context(), id)
	if err != nil {
		core_errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.UserResponseFromDomain(user))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core_errors.HandleError(c, fmt.Errorf(
			"convert id to int error: %w",
			core_errors.ErrInvalidArgument,
		))
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		core_errors.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
