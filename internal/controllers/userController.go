package controllers

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var validate = validator.New()
var customValidator = internal.NewValidator()

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

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Translate(nil)})
		return
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

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
		return
	}

	updatedUserDto, err := userService.UpdateUser(uint(userId), userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, updatedUserDto)

}
func UpdatePassword(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var passwordUpdateDto dtos.PasswordUpdateDto
	if err := c.ShouldBindJSON(&passwordUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = customValidator.Validator.Struct(passwordUpdateDto)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validatonError := range validationErrors {
			if validatonError.Tag() == "password-strength" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "The Password is not include one or more of the following characters : Uppercase , LowerCase , Special , Number "})
				return
			}
			return

		}
	}

	err = userService.UpdatePassword(uint(userId), passwordUpdateDto)
	if err != nil {
		if errors.Is(err, services.ErrIncorrectPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "The old password is incorrect"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update password. Please try again later."})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}
