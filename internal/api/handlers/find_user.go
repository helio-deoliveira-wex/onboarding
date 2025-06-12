package handlers

import (
	"net/http"

	"onboarding/internal/api/handlers/response"
	"onboarding/internal/service"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")
	u := service.GetUserById(id)

	if u == nil {
		c.IndentedJSON(http.StatusNotFound, response.ErrorResponse{Message: "User not found"})
	} else {
		c.IndentedJSON(http.StatusOK, u)
	}
}
