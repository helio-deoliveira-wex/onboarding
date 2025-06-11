package app_error

type ValidationError struct {
	OnboardingError
	Details []string
}

func (e ValidationError) Error() string {
	return e.OnboardingError.Message
}

func (e ValidationError) HasDetails() bool {
	return len(e.Details) > 0
}

func (e ValidationError) GetDetails() []string {
	return e.Details
}

func (e ValidationError) ErrorType() AppErrorType {
	return ValidationErrorType
}
