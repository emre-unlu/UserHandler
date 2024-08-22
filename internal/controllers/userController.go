package controllers

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/internal/utils"
	"github.com/emre-unlu/GinTest/pkg/customValidator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	userListDto := dtos.NewUserListDto()

	if p := c.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			userListDto.Page = parsedPage
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil {
			userListDto.Limit = parsedLimit
		}
	}

	users, total, err := userService.GetUserList(userListDto.Page, userListDto.Limit)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"total users": total, "users": users})
}

func GetUserById(c *gin.Context) {

	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
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
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Translate(nil)})
		return
	}

	createdUser, generatedPassword, err := userService.CreateUser(userDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": createdUser.ID, "password": generatedPassword})
}

func DeactivateUserById(c *gin.Context) {

	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
	}

	userDto, err := userService.DeactivateUserById(uint(userid))
	if err != nil {
		if errors.Is(err, internal.ErrUserAlreadyDisactive) {
			utils.RespondWithError(c, http.StatusConflict, err.Error())
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to deactivate user")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userDto, "message": "User successfully deactivated"})
}

func ActivateUserById(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
	}

	userDto, err := userService.ActivateUserById(uint(userid))
	if err != nil {
		if errors.Is(err, internal.ErrUserAlreadyActive) {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(err, internal.ErrUserDeleted) {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(c, http.StatusNotFound, "User not found")
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, "User activation failed")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userDto, "message": "User successfully reactivated"})
}

func SuspendUserById(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
	}

	userDto, err := userService.SuspendUserById(uint(userid))
	if err != nil {
		//To do fix these
		if errors.Is(err, internal.ErrUserAlreadySuspended) {
			utils.RespondWithError(c, http.StatusConflict, err.Error())
		} else if errors.Is(err, internal.ErrUserDeleted) {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondWithError(c, http.StatusNotFound, err.Error())
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to suspend user")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userDto, "message": "User successfully suspended"})
}

func UpdateUser(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User Id not found")
		return
	}
	userId, ok := id.(uint)
	if !ok {
		utils.RespondWithError(c, http.StatusUnauthorized, "Failed to parse user id")
		return
	}
	var userDto dtos.UserDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(userDto); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.RespondWithError(c, http.StatusBadRequest, validationErrors.Error())
		return
	}

	updatedUserDto, err := userService.UpdateUser(userId, userDto)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedUserDto)
}

func UpdatePassword(c *gin.Context) {
	id := c.Param("id")
	userid, err := strconv.Atoi(id)
	if err != nil || userid < 1 {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user id parameter")
	}

	var passwordUpdateDto dtos.PasswordUpdateDto
	if err := c.ShouldBindJSON(&passwordUpdateDto); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
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
		if errors.Is(err, internal.ErrIncorrectPassword) {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		} else {
			utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}
