package service

import (
	"onboarding/internal/model"
)

type IUserRepository interface {
	GetUser(id string) *model.User
	GetAllUsers() []model.User
	GetTotalUsers() int
	PutUser(u *model.User) error
}
