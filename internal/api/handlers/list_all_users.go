package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/service"
)

func ListAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAllUsers())
}
