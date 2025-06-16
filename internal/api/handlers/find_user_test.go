package handlers

import (
	"net/http"
	"net/http/httptest"
	"onboarding/internal/model"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock IUserService
type mockUserService struct {
	getUserByIdFunc func(string) *model.User
}

func (m *mockUserService) GetUserById(id string) *model.User {
	return m.getUserByIdFunc(id)
}

func (m *mockUserService) GetAllUsers() []model.User    { return nil }
func (m *mockUserService) SaveUser(u *model.User) error { return nil }

func TestFindUserHandler_UserFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockUserService{
		getUserByIdFunc: func(id string) *model.User {
			if id == "123" {
				return &model.User{Id: "123", FirstName: "Test"}
			}
			return nil
		},
	}
	handler := NewFindUserHandler((IUserService)(mockSvc))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	handler.FindUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Id":"123"`)
}

func TestFindUserHandler_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockUserService{
		getUserByIdFunc: func(id string) *model.User { return nil },
	}
	handler := NewFindUserHandler((*IUserService)(&mockSvc))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "notfound"}}

	handler.FindUser(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "User not found"))
}
