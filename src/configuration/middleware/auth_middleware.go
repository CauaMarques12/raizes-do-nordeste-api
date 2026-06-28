package middleware

import (
	"strings"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			err := rest_err.NewUnauthorizedRequestError("Authorization token is required")
			c.AbortWithStatusJSON(err.Code, err)
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := rest_err.NewUnauthorizedRequestError("Invalid authorization header")
			c.AbortWithStatusJSON(err.Code, err)
			return
		}

		claims, err := security.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(err.Code, err)
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("userRole")
		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		err := rest_err.NewForbiddenError("User does not have permission")
		c.AbortWithStatusJSON(err.Code, err)
	}
}
