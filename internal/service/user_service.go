package service

import (
	"math/rand"
	"strconv"

	"onboarding/internal/model"
	"onboarding/internal/util"
)

type UserService struct {
	repo      *IUserRepository
	validator *UserValidation
}

func NewUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: &repo, validator: NewUserValidation()}
}

func (us *UserService) GetUserById(id string) *model.User {
	return (*us.repo).GetUser(id)
}

func (us *UserService) GetAllUsers() []model.User {
	return (*us.repo).GetAllUsers()
}

func (us *UserService) SaveUser(u *model.User) error {
	if err := us.validator.Validate(u); err != nil {
		return err
	}
	// TODO validar regras de nomes únicos no BD

	return us.insertOrUpdate(u)
}

func (us *UserService) insertOrUpdate(u *model.User) error {
	if u.Id == "" {
		u.Id = "user-" + strconv.Itoa(1000+rand.Intn(1000))
	}

	if err := (*us.repo).PutUser(u); err != nil {
		util.AppLogger.Error("Error saving user",
			"user", u,
			"error", err)
		return err
	}
	return nil
}
