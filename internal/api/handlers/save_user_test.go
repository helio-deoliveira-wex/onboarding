package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"onboarding/internal/pkg/api_response"
	"onboarding/internal/service"
	"onboarding/internal/util"
	"onboarding/mocks"
	"testing"
)

func TestSaveUserHandler_SaveUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type testCase struct {
		name           string
		inputBody      string
		mockSetup      func(*mocks.IUserService)
		expectedStatus int
		bodyValidation func(*[]byte)
	}

	tests := []testCase{
		{
			name: "User saved successfully",
			inputBody: `{
							"id":"1",
							"first_name":"Test",
							"last_name":"User",
							"email":"user@test.com",
							"age":30
						}`,
			mockSetup: func(m *mocks.IUserService) {
				m.On("SaveUser", mock.AnythingOfType("*model.User")).Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Internal error while saving user",
			inputBody: `{
							"id":"1",
							"first_name":"Test",
							"last_name":"User",
							"email":"user@test.com",
							"age":30
						}`,
			mockSetup: func(m *mocks.IUserService) {
				m.On("SaveUser", mock.AnythingOfType("*model.User")).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			bodyValidation: func(b *[]byte) {
				errResponse := api_response.ErrorResponse{}
				err := json.Unmarshal(*b, &errResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errResponse.Message)
			},
		},
		{
			name:      "Invalid user data",
			inputBody: `{"id":"1"}`,
			mockSetup: func(m *mocks.IUserService) {
				validationErr := service.ErrUserValidation{
					Message: "validation failed", Details: []string{"Error details"}}
				m.On("SaveUser", mock.AnythingOfType("*model.User")).Return(validationErr)
			},
			expectedStatus: http.StatusBadRequest,
			bodyValidation: func(b *[]byte) {
				errResponse := api_response.ErrorResponse{}
				err := json.Unmarshal(*b, &errResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errResponse.Message)
				assert.NotNil(t, errResponse.Details)
				assert.True(t, len(errResponse.Details) > 0)
			},
		},
		{
			name:           "Invalid JSON request",
			inputBody:      ``,
			expectedStatus: http.StatusBadRequest,
			bodyValidation: func(b *[]byte) {
				errResponse := api_response.ErrorResponse{}
				err := json.Unmarshal(*b, &errResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errResponse.Message)
				assert.NotNil(t, errResponse.Details)
				assert.True(t, len(errResponse.Details) > 0)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.IUserService)
			if tc.mockSetup != nil {
				tc.mockSetup(mockService)
			}
			handler := NewSaveUserHandler(mockService)
			util.LoggerConfig()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(tc.inputBody))
			c.Request.Header.Set("Content-Type", "application/json")

			handler.SaveUser(c)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.bodyValidation != nil {
				body := w.Body.Bytes()
				tc.bodyValidation(&body)
			}
			mockService.AssertExpectations(t)
		})
	}
}
