package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"onboarding/internal/api/handlers/response"
	"onboarding/internal/database"
	"onboarding/internal/service"
	"onboarding/internal/service/app_error"
	"onboarding/internal/util"
)

func SaveUser(c *gin.Context) {
	var newUser database.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request"})
		return
	}

	u, err := service.SaveUser(newUser)
	if err != nil {
		util.AppLogger.Error("Error processing request to save user",
			"user", newUser,
			"error", err)

		cause := errors.Cause(err)
		util.AppLogger.Error(cause.Error())

		if appError, ok := cause.(app_error.AppError); ok &&
			appError.ErrorType() == app_error.ValidationErrorType {
			c.IndentedJSON(http.StatusBadRequest, response.Format(appError))
		} else {
			c.IndentedJSON(http.StatusInternalServerError, cause)
		}

		return
	}

	c.IndentedJSON(http.StatusCreated, u)
}
