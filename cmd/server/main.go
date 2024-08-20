package main

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/emre-unlu/GinTest/internal/routes"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/pkg/customValidator"
	"github.com/emre-unlu/GinTest/pkg/postgresql"
	"github.com/gin-gonic/gin"
)

func main() {
	postgresql.ConnectDatabase()
	userRepo := postgresql.NewPGUserRepository(postgresql.DB)
	userService := services.NewUserService(userRepo)
	customValidator := customValidator.NewValidator(userRepo)
	controllers.InitializeUserController(userService, customValidator)
	authService := services.NewAuthService(userRepo)
	controllers.InitilizeAuthController(authService)

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
