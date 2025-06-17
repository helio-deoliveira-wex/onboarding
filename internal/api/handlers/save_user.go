package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/model"
	"onboarding/internal/pkg/api_response"
	"onboarding/internal/service"
	"onboarding/internal/util"
)

type SaveUserHandler struct {
	userService *IUserService
}

func NewSaveUserHandler(us IUserService) *SaveUserHandler {
	return &SaveUserHandler{userService: &us}
}

func (h *SaveUserHandler) SaveUser(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, api_response.ErrorResponse{
			Message: "Invalid request", Details: []string{"Invalid JSON format"}})
		return
	}

	if err := (*h.userService).SaveUser(&newUser); err != nil {
		util.AppLogger.Error("Error processing request to save user",
			"user", newUser,
			"error", err)

		var errUserValidation = service.ErrUserValidation{}
		if errors.As(err, &errUserValidation) {
			c.IndentedJSON(http.StatusBadRequest, err)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}
