package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"onboarding/internal/model"
	"testing"
)

func TestUserValidation_Validate_ValidUser(t *testing.T) {
	t.Run("No error for valid user", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Age:       25,
		}
		err := uv.Validate(user)
		require.Nil(t, err)
	})
}

func TestUserValidation_Validate_InvalidAge(t *testing.T) {
	t.Run("Error for invalid age", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Age:       17,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, 1, len(errValidation.Details))
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, ErrUserInvalidAge.Error(), errValidation.Details[0])
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}

func TestUserValidation_Validate_MissingFirstName(t *testing.T) {
	t.Run("Error for missing first name", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "",
			LastName:  "Doe",
			Email:     "john@example.com",
			Age:       25,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, 1, len(errValidation.Details))
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, ErrUserInvalidName.Error(), errValidation.Details[0])
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}

func TestUserValidation_Validate_MissingLastName(t *testing.T) {
	t.Run("Error for missing last name", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "John",
			LastName:  "",
			Email:     "john@example.com",
			Age:       25,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, 1, len(errValidation.Details))
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, ErrUserInvalidName.Error(), errValidation.Details[0])
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}

func TestUserValidation_Validate_MissingEmail(t *testing.T) {
	t.Run("Error for missing email", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			Age:       25,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, 1, len(errValidation.Details))
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, ErrUserEmailRequired.Error(), errValidation.Details[0])
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}

func TestUserValidation_Validate_InvalidEmailFormat(t *testing.T) {
	t.Run("Error for invalid email format", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johnexample.com",
			Age:       25,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, 1, len(errValidation.Details))
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, ErrUserEmailInvalidFormat.Error(), errValidation.Details[0])
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}

func TestUserValidation_Validate_MultipleErrors(t *testing.T) {
	t.Run("Multiple errors for invalid user", func(t *testing.T) {
		uv := NewUserValidation()
		user := &model.User{
			FirstName: "",
			LastName:  "",
			Email:     "",
			Age:       10,
		}
		err := uv.Validate(user)
		require.Error(t, err)
		if errValidation, ok := err.(ErrUserValidation); ok {
			assert.Equal(t, InvalidUserMessage, errValidation.Message)
			assert.Equal(t, 3, len(errValidation.Details))
		} else {
			assert.FailNow(t, "Expected ErrUserValidation type, got %T", err)
		}
	})
}
