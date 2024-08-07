package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.CreateUser)
	router.POST("/users/:id", controllers.DeleteUserById)
}
