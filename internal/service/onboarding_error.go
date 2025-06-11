package service

import "encoding/json"

type ErrorType string

const (
	OnboardingError ErrorType = "OnboardingError"
	ValidationError ErrorType = "ValidationError"
)

type AppError struct {
	Message   string   `json:"error"`
	Details   []string `json:"details"`
	ErrorType ErrorType
}

func (e *AppError) Error() string {
	str, _ := json.Marshal(e)
	return string(str)
}
