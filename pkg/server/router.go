package server

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/handler"
	"github.com/todo_manager/pkg/middleware"
)

func RegisterRouters(router *gin.Engine, handler handler.HandlerInterface, mid middleware.MidInterface) {

	userRelatedTasks := router.Group("/user")
	{
		userRelatedTasks.POST("/signup", handler.SignUp)
		userRelatedTasks.GET("/signin", handler.SignIn)
		userRelatedTasks.GET("/viewallusers", mid.Middleware(), handler.ViewAllUsers)
		userRelatedTasks.GET("/viewuser", handler.ViewUser)
		userRelatedTasks.PUT("/updatealldetails/:id", handler.UpdateAllDetails)
		userRelatedTasks.PATCH("/updatedetail", handler.UpdateDetail)
		userRelatedTasks.DELETE("/deleteuser/:id", handler.DeleteUser)
	}

	taskRelatedTasks := router.Group("/tasks")
	{
		taskRelatedTasks.POST("/createtask/:id", handler.CreateTask)
		taskRelatedTasks.GET("/viewalltask", handler.ViewAllTask)
		taskRelatedTasks.GET("/viewaalltaskofuser/:id", handler.ViewAllTaskOfUser)
		taskRelatedTasks.GET("/viewtask/:id", handler.ViewTask)
		taskRelatedTasks.PATCH("/updatetaskstatus", handler.UpdateTaskStatus)
		taskRelatedTasks.PUT("/updatealldetailsoftask/:id", handler.UpadteAllTaskDetail)
		taskRelatedTasks.DELETE("/deletetask/:id", handler.DeleteTask)
	}
}
