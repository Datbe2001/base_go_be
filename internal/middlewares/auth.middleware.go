package middlewares

import (
	"base_go_be/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.ErrorResponse(c, 401, response.ErrInvalidToken)
			c.Abort()
			return
		}
		c.Next()
	}
}
