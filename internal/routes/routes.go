package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id/deactivate", controllers.DeactivateUserById)
	router.PUT("/users/:id/suspend", controllers.SuspendUserById)
	router.PUT("/users/:id/activate", controllers.ActivateUserById)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.PUT("/users/:id/password", controllers.UpdatePassword)
}
