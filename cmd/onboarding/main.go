package main

import (
	"github.com/gin-gonic/gin"
	"onboarding/internal/api"
	"onboarding/internal/util"
)

func main() {
	//Logs:
	util.LoggerConfig()

	//HTTP server:
	router := gin.Default()
	api.ConfigRoutes(router)
	router.Run("localhost:8080")
}
