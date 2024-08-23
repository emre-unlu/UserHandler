package utils

import (
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
	errorResponse := dtos.NewErrorDto(message, code)
	c.JSON(code, errorResponse)
	c.Abort()

}
