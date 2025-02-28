package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Mid) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "autharization header is empty",
			})
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "token string cannot be validated",
			})
			return
		}
		claims, err := m.auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "token is not validated ",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"Message": "validated successfully",
		})
		c.Set("claims", claims)
		c.Next()
	}
}
