package middlewares

import (
	"fmt"
	"strings"

	"github.com/Talan-Application/talan-back/internal/core/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := fmt.Errorf("authorization header is required: %w", core_errors.ErrUnauthorized)
			core_errors.HandleError(c, err)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			err := fmt.Errorf("authorization header format must be Bearer {token}: %w", core_errors.ErrUnauthorized)
			core_errors.HandleError(c, err)
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			err := fmt.Errorf("invalid or expired token: %w", core_errors.ErrUnauthorized)
			core_errors.HandleError(c, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(float64); ok {
				c.Set("userId", int(sub))
			} else {
				c.Set("userId", claims["sub"])
			}
		}

		c.Next()
	}
}
