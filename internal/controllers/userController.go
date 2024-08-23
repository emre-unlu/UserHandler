package controllers

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/internal/utils"
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
	customValidate = customValidator
}

func GetUserList(c *gin.Context) {
	userListRequestDto := dtos.UserListRequestDto{}
	if err := c.ShouldBindQuery(&userListRequestDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validate.Struct(userListRequestDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Translate(nil)})
		return
	}

	users, err := userService.GetUserList(userListRequestDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {

	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
		return
	}

	user, err := userService.GetUserById(uint(userid))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var userDto dtos.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Translate(nil)})
		return
	}

	createdUser, err := userService.CreateUser(userDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": createdUser.ID})
}

func SuspendUserById(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)

	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
		return
	}

	err = userService.SuspendUserById(uint(userid))

	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User with ID: %d successfully suspended", userid),
	})
}

func DeactivateUserById(c *gin.Context) {

	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
		return
	}

	err = userService.DeactivateUserById(uint(userid))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID: %d successfully deactivated", userid)})
}

func ActivateUserById(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)

	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
		return
	}

	err = userService.ActivateUserById(uint(userid))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID: %d successfully reactivated", userid)})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)

	var userDto dtos.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.RespondWithError(c, http.StatusBadRequest, validationErrors.Error())
		return
	}

	err = userService.UpdateUser(uint(userid), userDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID: %d successfully updated with the given data", userid)})
}

func UpdatePassword(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)

	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
		return
	}

	var passwordUpdateDto dtos.PasswordUpdateDto
	if err := c.ShouldBindJSON(&passwordUpdateDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := customValidate.Validator.Struct(passwordUpdateDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		translatedErrors := validationErrors.Translate(customValidate.Translator)
		c.JSON(http.StatusBadRequest, gin.H{"errors": translatedErrors})
		return
	}

	err = userService.UpdatePassword(uint(userid), passwordUpdateDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Password of user with Id : %d successfully updated ", userid)})
}

func ForgotPassword(c *gin.Context) {
	forgotPasswordDto := dtos.ForgotPasswordDto{}
	if err := c.ShouldBindJSON(&forgotPasswordDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := customValidate.Validator.Struct(forgotPasswordDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		translatedErrors := validationErrors.Translate(customValidate.Translator)
		c.JSON(http.StatusBadRequest, gin.H{"errors": translatedErrors})
		return
	}

	err := userService.ForgotPassword(forgotPasswordDto.Email)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Email with new password is successfully sent to the email : %s ", forgotPasswordDto.Email)})

}
