package middlewares

import (
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/gin-gonic/gin"
)

func GrantAccess(allowedRoles ...domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("currentUser").(domain.User)

		if !user.HasAnyRole(allowedRoles...) {
			c.AbortWithStatusJSON(403, gin.H{"error": "access denied"})
			return
		}
		c.Next()
	}
}
