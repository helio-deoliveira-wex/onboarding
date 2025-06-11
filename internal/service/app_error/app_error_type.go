package app_error

type AppErrorType string

const (
	OnboardingErrorType AppErrorType = "OnboardingError"
	ValidationErrorType AppErrorType = "ValidationError"
)
