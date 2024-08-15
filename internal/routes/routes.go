package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUserById)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.PUT("/users/:id/password", controllers.UpdatePassword)
}
