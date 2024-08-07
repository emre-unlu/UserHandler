package controllers

import (
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userService *services.UserService

func InitializeUserController(service *services.UserService) {
	userService = service
}

func GetUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.DeleteUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
