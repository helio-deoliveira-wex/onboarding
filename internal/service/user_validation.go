package service

import (
	"errors"
	"onboarding/internal/model"
	"strings"
)

const InvalidUserMessage = "User did not pass validation"

var (
	ErrUserInvalidAge         = errors.New("User does not meet minimum age requirement")
	ErrUserInvalidName        = errors.New("User first/last names required")
	ErrUserEmailRequired      = errors.New("User email required")
	ErrUserEmailInvalidFormat = errors.New("User email must be properly formatted")
)

type ErrUserValidation struct {
	Message string   `json:"error"`
	Details []string `json:"details,omitempty"`
}

func (e ErrUserValidation) Error() string {
	return InvalidUserMessage
}

func (e ErrUserValidation) Is(target error) bool {
	return target != nil && target.Error() == e.Error()
}

type UserValidation struct {
}

func NewUserValidation() *UserValidation {
	return &UserValidation{}
}

func (v *UserValidation) Validate(u *model.User) *ErrUserValidation {
	errList := make([]string, 0, 10)
	if err := v.validateAge(u.Age); err != nil {
		errList = append(errList, err.Error())
	}
	if err := v.validateName(u.FirstName); err != nil {
		errList = append(errList, err.Error())
	} else if err := v.validateName(u.LastName); err != nil {
		errList = append(errList, err.Error())
	}
	if err := v.validateEmail(u.Email); err != nil {
		errList = append(errList, err.Error())
	}

	if len(errList) > 0 {
		return &ErrUserValidation{Message: InvalidUserMessage, Details: errList}
	}
	return nil
}

func (v *UserValidation) validateAge(a int) error {
	if a < 18 {
		return ErrUserInvalidAge
	}
	return nil
}

func (v *UserValidation) validateName(n string) error {
	if v.isEmptyName(n) {
		return ErrUserInvalidName
	}
	return nil
}

func (v *UserValidation) validateEmail(e string) error {
	if v.isEmptyName(e) {
		return ErrUserEmailRequired
	}
	if !strings.Contains(e, "@") {
		return ErrUserEmailInvalidFormat
	}
	return nil
}

func (v *UserValidation) isEmptyName(n string) bool {
	n = strings.TrimSpace(n)
	return len(n) == 0
}
