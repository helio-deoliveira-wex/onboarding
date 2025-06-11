package app_error

type OnboardingError struct {
	Message string
}

func (e OnboardingError) Error() string {
	return e.Message
}

func (e OnboardingError) ErrorType() AppErrorType {
	return OnboardingErrorType
}
