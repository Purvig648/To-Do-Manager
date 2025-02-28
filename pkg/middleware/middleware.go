package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/auth"
)

type Mid struct {
	auth auth.Authentication
}

type MidInterface interface {
	Middleware() gin.HandlerFunc
}

func NewMiddleware(a auth.Authentication) MidInterface {
	return &Mid{
		auth: a,
	}
}
