package middlewares

import (
	"base_go_be/pkg/config"
	"base_go_be/pkg/jwt"
	"base_go_be/pkg/response"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ErrorResponse(c, 401, response.ErrInvalidToken)
			c.Abort()
			return
		}

		// Check if the header has the "Bearer" prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.ErrorResponse(c, 401, response.ErrInvalidToken)
			c.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateToken(tokenString, config.JWT.SecretKey)
		if err != nil {
			if errors.Is(err, jwt.ErrExpiredToken) {
				response.ErrorResponse(c, 401, "Token expired")
			} else {
				response.ErrorResponse(c, 401, response.ErrInvalidToken)
			}
			c.Abort()
			return
		}

		// Store user info in context for later use
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleMiddleware checks if the user has a required role
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			response.ErrorResponse(c, 401, response.ErrInvalidToken)
			c.Abort()
			return
		}

		// Check if a user role is in the allowed roles
		hasRole := false
		for _, role := range roles {
			if userRole == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			response.ErrorResponse(c, 403, "Forbidden: insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}
