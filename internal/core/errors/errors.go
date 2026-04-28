package core_errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
)

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// HandleError maps domain errors to HTTP responses and sends them via Gin
func HandleError(c *gin.Context, err error) {
	var appErr *AppError

	switch {
	case errors.Is(err, ErrNotFound):
		appErr = &AppError{Code: http.StatusNotFound, Message: err.Error()}
	case errors.Is(err, ErrConflict):
		appErr = &AppError{Code: http.StatusConflict, Message: err.Error()}
	case errors.Is(err, ErrInvalidArgument):
		appErr = &AppError{Code: http.StatusBadRequest, Message: err.Error()}
	default:
		appErr = &AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}

	c.AbortWithStatusJSON(appErr.Code, appErr)
}
