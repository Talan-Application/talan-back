package core_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
)

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError

	switch {
	case errors.Is(err, ErrNotFound):
		appErr = &AppError{Code: http.StatusNotFound, Message: err.Error()}
	case errors.Is(err, ErrConflict):
		appErr = &AppError{Code: http.StatusConflict, Message: err.Error()}
	case errors.Is(err, ErrInvalidArgument):
		appErr = &AppError{Code: http.StatusBadRequest, Message: err.Error()}
	case errors.Is(err, ErrUnauthorized):
		appErr = &AppError{Code: http.StatusUnauthorized, Message: err.Error()}
	case errors.Is(err, ErrForbidden):
		appErr = &AppError{Code: http.StatusForbidden, Message: err.Error()}
	default:
		fmt.Println("internal error: ", err)
		appErr = &AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}

	c.AbortWithStatusJSON(appErr.Code, appErr)
}
