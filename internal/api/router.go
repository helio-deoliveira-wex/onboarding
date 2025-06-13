package api

import (
	"github.com/gin-gonic/gin"
	"onboarding/internal/api/handlers"
)

var (
	saveUserHandler     *handlers.SaveUserHandler
	findUserHandler     *handlers.FindUserHandler
	listAllUsersHandler *handlers.ListAllUsersHandler
)

func ConfigRoutes(router *gin.Engine, us *handlers.IUserService) {
	saveUserHandler = handlers.NewSaveUserHandler(us)
	findUserHandler = handlers.NewFindUserHandler(us)
	listAllUsersHandler = handlers.NewListAllUsersHandler(us)

	router.POST("/save", saveUser)
	router.GET("/find/:id", findUser)
	router.GET("/list", listAllUsers)
}

func saveUser(c *gin.Context) {
	saveUserHandler.SaveUser(c)
}

func findUser(c *gin.Context) {
	findUserHandler.FindUser(c)
}

func listAllUsers(c *gin.Context) {
	listAllUsersHandler.ListAllUsers(c)
}
