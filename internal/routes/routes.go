package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/emre-unlu/GinTest/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/refresh", controllers.RefreshToken)
	router.POST("/users", controllers.CreateUser)

	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/users", controllers.GetUsersWithPagination)
		authorized.GET("/users/getUser", controllers.GetUserById)
		authorized.PUT("/users/:id/deactivate", controllers.DeactivateUserById)
		authorized.PUT("/users/:id/suspend", controllers.SuspendUserById)
		authorized.PUT("/users/:id/activate", controllers.ActivateUserById)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.PUT("/users/:id/password", controllers.UpdatePassword)
	}
}
