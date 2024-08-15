package controllers

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/pkg/customValidator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var validate = validator.New()
var customValidate *customValidator.CustomValidator
var userService *services.UserService

func InitializeUserController(service *services.UserService, customValidator *customValidator.CustomValidator) {
	userService = service
	customValidate = customValidator // Assign custom validator to global variable
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

	if err := customValidate.Validator.Struct(passwordUpdateDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		translatedErrors := validationErrors.Translate(customValidate.Translator)
		c.JSON(http.StatusBadRequest, gin.H{"errors": translatedErrors})
		return
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
