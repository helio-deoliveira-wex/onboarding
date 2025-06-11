package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/service"
)

func FindUser(c *gin.Context) {
	id := c.Param("id")
	u := service.GetUserById(id)

	if u == nil {
		c.IndentedJSON(http.StatusNotFound, &service.AppError{
			Message: "User not found", ErrorType: service.OnboardingError})
	} else {
		c.IndentedJSON(http.StatusOK, u)
	}
}
