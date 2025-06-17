package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"onboarding/internal/model"
	"onboarding/internal/util"
	"onboarding/mocks"
	"testing"
	"time"
)

func Test_GetUserById(t *testing.T) {
	t.Run("User exists", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		u := createUser("1")
		mockRepo.On("GetUser", "1").Return(&u)
		userService := NewUserService(mockRepo)

		user := userService.GetUserById("1")

		assert.NotNil(t, user)
		assert.Equal(t, u, *user)
		mockRepo.AssertExpectations(t)
	})
	t.Run("User does not exist", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		mockRepo.On("GetUser", "1").Return(nil)
		userService := NewUserService(mockRepo)

		user := userService.GetUserById("1")

		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("Return all users", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		list := []model.User{
			createUser("1"),
			createUser("2"),
		}
		mockRepo.On("GetAllUsers").Return(list)
		userService := NewUserService(mockRepo)

		users := userService.GetAllUsers()

		assert.NotNil(t, users)
		assert.Len(t, users, 2)
		assert.Equal(t, list, users)
		mockRepo.AssertExpectations(t)

	})
	t.Run("There is no user", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		list := []model.User{}
		mockRepo.On("GetAllUsers").Return(list)
		userService := NewUserService(mockRepo)

		users := userService.GetAllUsers()

		assert.NotNil(t, users)
		assert.Equal(t, list, users)
		mockRepo.AssertExpectations(t)

	})
}

func TestSaveUser(t *testing.T) {
	t.Run("Validation error when saving user", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		userService := NewUserService(mockRepo)

		err := userService.SaveUser(&model.User{Id: "2"})

		assert.NotNil(t, err)
		assert.ErrorAs(t, err, &ErrUserValidation{})
		assert.True(t, len(err.(ErrUserValidation).Details) > 0)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Repository error when saving user", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		e := errors.New("Test error")
		mockRepo.On("PutUser", mock.AnythingOfType("*model.User")).Return(e)
		userService := NewUserService(mockRepo)
		u := createUser("1")
		util.LoggerConfig()

		err := userService.SaveUser(&u)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, e)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Saving user without ID", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		mockRepo.On("PutUser", mock.AnythingOfType("*model.User")).Return(nil)
		userService := NewUserService(mockRepo)
		u := createUser("")

		err := userService.SaveUser(&u)

		assert.Nil(t, err)
		assert.NotNil(t, u.Id)
		assert.True(t, len(u.Id) > 0, "Expected user ID to be set")
		mockRepo.AssertExpectations(t)
	})
	t.Run("Saving user with ID", func(t *testing.T) {
		mockRepo := mocks.NewIUserRepository(t)
		mockRepo.On("PutUser", mock.AnythingOfType("*model.User")).Return(nil)
		userService := NewUserService(mockRepo)
		u := createUser("3")

		err := userService.SaveUser(&u)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func createUser(id string) model.User {
	return model.User{
		Id:        id,
		FirstName: "Test-" + id,
		LastName:  "User",
		Email:     "test.user@test.com",
		Age:       30,
		UpdatedAt: time.Now(),
	}
}
