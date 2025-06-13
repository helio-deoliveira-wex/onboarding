package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onboarding/internal/pkg/api_response"
)

type FindUserHandler struct {
	userService *IUserService
}

func NewFindUserHandler(us *IUserService) *FindUserHandler {
	return &FindUserHandler{userService: us}
}

func (fu *FindUserHandler) FindUser(c *gin.Context) {
	id := c.Param("id")
	u := (*fu.userService).GetUserById(id)

	if u == nil {
		c.IndentedJSON(http.StatusNotFound, api_response.ErrorResponse{Message: "User not found"})
	} else {
		c.IndentedJSON(http.StatusOK, u)
	}
}
