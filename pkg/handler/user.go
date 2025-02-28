package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todo_manager/pkg/auth"
	"github.com/todo_manager/pkg/model"
	"github.com/todo_manager/pkg/util"
)

func (h *handler) SignUp(c *gin.Context) {
	signUpData := model.SignUp{}
	err := c.BindJSON(&signUpData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	userID, statusCode, err := h.svc.SignUpService(signUpData)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"error": nil,
		"response": gin.H{
			"user_id": userID,
			"message": "created user successfully",
		},
	})
}

func (h *handler) SignIn(c *gin.Context) {
	signInData := model.SignIn{}
	err := c.BindJSON(&signInData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	token, statusCode, err := h.svc.SignInService(signInData)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": "signIn Failed",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"error":    nil,
		"response": token,
		"Message":  "sign in successful",
	})
}

func (h *handler) ViewAllUsers(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userClaims := claims.(*auth.Claims)
	c.JSON(http.StatusOK, gin.H{
		"message": "Viewing all users",
		"user":    userClaims.ID,
	})
	resp, statusCode, err := h.svc.ViewAllUsers()
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": "no data displayed",
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
	})
}

func (h *handler) ViewUser(c *gin.Context) {
	userDetail := model.UserRequest{}
	err := c.BindJSON(&userDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	resp, statusCode, err := h.svc.ViewUser(userDetail)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": "no data displayed",
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
	})
}

func (h *handler) UpdateAllDetails(c *gin.Context) {
	id := c.Param("id")
	userDetail := model.UserDetailsUpdate{}
	err := c.BindJSON(&userDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "could not bindJSON",
		})
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "could not convert to int",
		})
		return
	}
	resp, statusCode, err := h.svc.UpdateAllDetails(uint(ID), userDetail)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": "could not update details",
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
		"Message":  "updated successfully",
	})
}

func (h *handler) UpdateDetail(c *gin.Context) {
	id := c.Query("id")
	choice := c.Query("choice")
	if _, ok := util.Choices[choice]; !ok {
		return
	}
	userDetail := model.UserDetailUpdate{}
	err := c.BindJSON(&userDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
		return
	}
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "could not convert to int",
		})
		return
	}
	resp, statusCode, err := h.svc.UpdateDetail(uint(uid), userDetail, choice)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": "could not update details",
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"error":    nil,
		"response": resp,
		"Message":  "updated successfully",
	})
}

func (h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
	}
	statusCode, err := h.svc.DeleteUser(uint(uid))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error":    err.Error(),
			"response": nil,
		})
	}
	c.JSON(statusCode, gin.H{
		"error":   nil,
		"Message": "deleted succesfully",
	})
}
