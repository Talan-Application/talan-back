package middlewares

import (
	"github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			core_errors.HandleError(c, err)
		}
	}
}
