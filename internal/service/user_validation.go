package service

import (
	"errors"
	"onboarding/internal/database"
	"onboarding/internal/service/app_error"
	"strings"
)

func validate(u database.User) *app_error.ValidationError {
	errList := make([]string, 0, 10)
	if err := validateAge(u.Age); err != nil {
		errList = append(errList, err.Error())
	}
	if err := validateName(u.FirstName); err != nil {
		errList = append(errList, err.Error())
	} else if err := validateName(u.LastName); err != nil {
		errList = append(errList, err.Error())
	}
	if err := validateEmail(u.Email); err != nil {
		errList = append(errList, err.Error())
	}

	if len(errList) > 0 {
		return &app_error.ValidationError{
			OnboardingError: app_error.OnboardingError{Message: "User did not pass validation"},
			Details:         errList}
	}
	return nil
}

func validateAge(a int) error {
	if a < 18 {
		return errors.New("User does not meet minimum age requirement")
	}
	return nil
}

func validateName(n string) error {
	if isEmptyName(n) {
		return errors.New("User first/last names required")
	}
	return nil
}

func validateEmail(e string) error {
	if isEmptyName(e) {
		return errors.New("User email required")
	}
	if !strings.Contains(e, "@") {
		return errors.New("User email must be properly formatted")
	}
	return nil
}

func isEmptyName(n string) bool {
	n = strings.TrimSpace(n)
	return len(n) == 0
}
