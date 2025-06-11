package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/api/handlers/response"
	"onboarding/internal/service"
)

func FindUser(c *gin.Context) {
	id := c.Param("id")
	u := service.GetUserById(id)

	if u == nil {
		c.IndentedJSON(http.StatusNotFound, response.ErrorResponse{Message: "User not found"})
	} else {
		c.IndentedJSON(http.StatusOK, u)
	}
}
