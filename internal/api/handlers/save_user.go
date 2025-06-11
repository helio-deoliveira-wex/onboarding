package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/database"
	"onboarding/internal/service"
	"onboarding/internal/util"
)

func SaveUser(c *gin.Context) {
	var newUser database.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, &service.AppError{
			Message: "Invalid request", ErrorType: service.OnboardingError})
		return
	}

	u, err := service.SaveUser(newUser)
	if err != nil {
		util.AppLogger.Error("Error saving user",
			"user", newUser,
			"errorMsg", err.Message,
			"errorType", err.ErrorType,
			"errorDetails", err.Details)
		statusCode := http.StatusInternalServerError //default error code
		if err.ErrorType == service.ValidationError {
			statusCode = http.StatusBadRequest
		}
		c.IndentedJSON(statusCode, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, u)
}
