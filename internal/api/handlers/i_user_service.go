package handlers

import (
	"onboarding/internal/model"
)

type IUserService interface {
	GetUserById(id string) *model.User
	GetAllUsers() []model.User
	SaveUser(u *model.User) error
}
