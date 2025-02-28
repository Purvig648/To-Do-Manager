package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Mid) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			return
		}
		claims, err := m.auth.ValidateToken(tokenString)
		if err != nil {
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
