package api

import (
	"github.com/gin-gonic/gin"
	"onboarding/internal/api/handlers"
)

func ConfigRoutes(router *gin.Engine, handler ) {
	router.POST("/save", handlers.SaveUser)
	router.GET("/find/:id", handlers.FindUser)
	router.GET("/list", handlers.ListAllUsers)
}
