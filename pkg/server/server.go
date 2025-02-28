package server

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/handler"
	"github.com/todo_manager/pkg/middleware"
)

func StartApp(handler handler.HandlerInterface, mid middleware.MidInterface) {
	router := gin.Default()

	RegisterRouters(router, handler, mid)

	router.Run(":8081")
}
