package middlewares

import (
	"github.com/Talan-Application/talan-back/internal/domain"
	"github.com/gin-gonic/gin"
)

func GrantAccess(allowedRoles ...domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("userRole")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "authentication required"})
			return
		}

		// Cast the interface to the UserRole type
		userRole, ok := val.(domain.UserRole)
		if !ok {
			c.AbortWithStatusJSON(500, gin.H{"error": "internal role error"})
			return
		}

		for _, r := range allowedRoles {
			if r == userRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"error": "access denied: insufficient permissions"})
	}
}
