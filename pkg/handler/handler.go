package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/service"
)

type handler struct {
	svc service.ServiceInterface
}

type HandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	ViewAllUsers(c *gin.Context)
	ViewUser(c *gin.Context)
	UpdateAllDetails(c *gin.Context)
	UpdateDetail(c *gin.Context)

	CreateTask(c *gin.Context)
	ViewAllTask(c *gin.Context)
	ViewAllTaskOfUser(c *gin.Context)
	ViewTask(c *gin.Context)
	UpdateTaskStatus(c *gin.Context)
	UpadteAllTaskDetail(c *gin.Context)
}

func NewHandlerLayer(svc service.ServiceInterface) HandlerInterface {
	return &handler{
		svc: svc,
	}
}
