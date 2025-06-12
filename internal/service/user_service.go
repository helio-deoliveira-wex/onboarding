package service

import (
	"math/rand"
	"strconv"

	"onboarding/internal/database"
	"onboarding/internal/model"
	"onboarding/internal/util"

	"github.com/pkg/errors"
)

type UserInterface interface {
	GetUser(id string) *model.User
	GetAllUsers() []*model.User
	GetTotalUsers() int
	PutUser(user model.User) error
}

func GetUserById(id string) *database.User {
	return database.GetUser(id)
}

func GetAllUsers() []database.User {
	return database.GetAllUsers()
}

func SaveUser(u database.User) (*database.User, error) {
	if e := validate(u); e != nil {
		return nil, errors.Wrap(e, "Error saving user")
	}
	// TODO validar regras de nomes únicos no BD

	if database.GetUser(u.Id) != nil {
		updateUser(u)
	} else {
		u = insertUser(u)
	}

	return &u, nil
}

func updateUser(u database.User) {
	database.PutUser(&u)
}

func insertUser(u database.User) database.User {
	if u.Id == "" {
		u.Id = "user-" + strconv.Itoa(1000+rand.Intn(1000))
	}
	database.PutUser(&u)
	util.AppLogger.Info("New user created",
		"user", u,
		"totalOfUsers", database.GetTotalUsers())
	return u
}
