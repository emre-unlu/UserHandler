package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/emre-unlu/GinTest/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/token/refresh", controllers.RefreshToken)
	router.POST("/users", controllers.CreateUser)

	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/users", controllers.GetUserList)
		authorized.GET("/users/:id", controllers.GetUserById)
		authorized.DELETE("/users/:id", controllers.DeactivateUserById)
		authorized.PATCH("/users/:id/suspend", controllers.SuspendUserById)
		authorized.PATCH("/users/:id/activate", controllers.ActivateUserById)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.PUT("/users/:id/password", controllers.UpdatePassword)
	}
}
