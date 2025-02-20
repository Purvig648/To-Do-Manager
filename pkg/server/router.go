package server

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/handler"
)

func RegisterRouters(router *gin.Engine, handler handler.HandlerInterface) {
	router.POST("/signup", handler.SignUp)
	router.GET("/signin", handler.SignIn)
	router.GET("/viewallusers", handler.ViewAllUsers)
	router.GET("/viewuser", handler.ViewUser)
	router.PUT("/updatealldetails/:id", handler.UpdateAllDetails)
	router.PATCH("/updatedetail", handler.UpdateDetail)
}
