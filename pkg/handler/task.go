package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/model"
)

func (h *handler) CreateTask(c *gin.Context) {
	id := c.Param("id")
	taskData := model.Task{}
	err := c.BindJSON(&taskData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	tid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}

	resp, statusCode, err := h.svc.CreateTask(uint(tid), taskData)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	taskResponse := model.TaskResponse{
		ID:              resp.ID,
		TaskName:        resp.TaskName,
		TaskDescription: resp.TaskDescription,
		TaskDeadline:    resp.TaskDeadline,
		TaskStatus:      resp.TaskStatus,
		UserID:          resp.UserID,
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": taskResponse,
	})
}

func (h *handler) ViewAllTask(c *gin.Context) {
	resp, statusCode, err := h.svc.ViewAllTask()
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
	})

}

func (h *handler) ViewAllTaskOfUser(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}

	resp, statusCode, err := h.svc.ViewAllTaskOfUser(uint(uid))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
	})
}
