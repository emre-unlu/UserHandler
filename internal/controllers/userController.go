package controllers

import (
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	var userDto dtos.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userDto.Name == "" {
		c.JSON(http.StatusBadRequest, "name is a required field")
		return
	}
	if userDto.Birthdate != "" {
		if _, err := time.Parse("2006-01-02", userDto.Birthdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birthdate format. Use YYYY-MM-DD"})
			return
		}
	}
	createdUser, generatedPassword, err := userService.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": createdUser, "password": generatedPassword})
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDto, err := userService.DeleteUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userDto dtos.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if userDto.Birthdate != "" {
		if _, err := time.Parse("2006-01-02", userDto.Birthdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birthdate format. Use YYYY-MM-DD"})
			return
		}
	}
	var newPassword string
	if password, exists := c.GetPostForm("password"); exists {
		newPassword = password
	}
	updatedUserDto, err := userService.UpdateUser(uint(userId), userDto, newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, updatedUserDto)

}
