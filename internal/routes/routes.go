package routes

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/emre-unlu/GinTest/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/token/refresh", controllers.RefreshToken)
	router.POST("/password/reset", controllers.ForgotPassword)

	authorized := router.Group("/users")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.POST("/", controllers.CreateUser)
		authorized.GET("/", controllers.GetUserList)
		authorized.GET("/:id", controllers.GetUserById)
		authorized.DELETE("/:id", controllers.DeactivateUserById)
		authorized.PUT("/:id/suspend", controllers.SuspendUserById)
		authorized.PUT("/:id/activate", controllers.ActivateUserById)
		authorized.PUT("/:id", controllers.UpdateUser)
		authorized.PUT("/:id/password", controllers.UpdatePassword)
	}

	profile := router.Group("/profile")
	profile.Use(middleware.JWTAuthMiddleware())
	{
		profile.GET("/", controllers.GetUserList)

	}
}
