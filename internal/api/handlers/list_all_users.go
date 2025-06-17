package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListAllUsersHandler struct {
	userService *IUserService
}

func NewListAllUsersHandler(us IUserService) *ListAllUsersHandler {
	return &ListAllUsersHandler{userService: &us}
}

func (lu *ListAllUsersHandler) ListAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, (*lu.userService).GetAllUsers())
}
