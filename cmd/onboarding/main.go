package main

import (
	"github.com/gin-gonic/gin"
	"onboarding/internal/api"
	"onboarding/internal/api/handlers"
	"onboarding/internal/database"
	"onboarding/internal/service"
	"onboarding/internal/util"
)

func main() {
	//Logs:
	util.LoggerConfig()

	//Dependency injection:
	var userRepository service.IUserRepository = database.NewUserRepo()
	var userService handlers.IUserService = service.NewUserService(&userRepository)

	//HTTP server:
	router := gin.Default()
	api.ConfigRoutes(router, &userService)
	router.Run("localhost:8080")
}
