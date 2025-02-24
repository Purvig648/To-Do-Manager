package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/model"
	"github.com/todo_manager/pkg/util"
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

func (h *handler) ViewTask(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	taskDetail, statusCode, err := h.svc.ViewTask(uint(tid))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": taskDetail,
	})
}

func (h *handler) UpdateTaskStatus(c *gin.Context) {
	id := c.Query("id")
	taskStatusChoice := c.Query("choice")
	if _, ok := util.TaskChoices[taskStatusChoice]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "not a valid request",
			"response": nil,
		})
		return
	}
	tid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
			"Message":  "could not convert",
		})
		return
	}
	resp, statusCode, err := h.svc.UpdateTaskStatus(uint(tid), taskStatusChoice)
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
		"Message":  "updated successfully",
	})
}

func (h *handler) UpadteAllTaskDetail(c *gin.Context) {
	id := c.Param("id")
	tid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
			"message":  "not a valid format",
		})
		return
	}
	taskUpdateDetails := model.TaskDetailsUpdate{}
	err = c.BindJSON(&taskUpdateDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
			"message":  "could not bind JSON",
		})
		return
	}
	resp, statusCode, err := h.svc.UpadteAllTaskDetail(uint(tid), taskUpdateDetails)
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
		"message":  "updated successfully",
	})
}
