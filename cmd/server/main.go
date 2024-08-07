package server

import (
	"github.com/emre-unlu/GinTest/internal/controllers"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/routes"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	userRepo := models.NewPGUserRepository(models.DB)
	userService := services.NewUserService(userRepo)
	controllers.InitializeUserController(userService)

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
