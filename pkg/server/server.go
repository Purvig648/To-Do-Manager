package server

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/handler"
)

func StartApp(handler handler.HandlerInterface) {
	router := gin.Default()

	RegisterRouters(router, handler)

	router.Run(":8081")
}
