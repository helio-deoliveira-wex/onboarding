package main

import (
	"onboarding/internal/api"
	"onboarding/internal/api/handlers"
	"onboarding/internal/database"
	"onboarding/internal/service"
	"onboarding/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	//Logs:
	util.LoggerConfig()

	//HTTP server:
	router := gin.Default()

	userRepo := database.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	
	api.ConfigRoutes(router, userRepo)
	router.Run("localhost:8080")
}
