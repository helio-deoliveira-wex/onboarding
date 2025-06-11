package service

import (
	"math/rand"
	"onboarding/internal/database"
	"onboarding/internal/util"
	"strconv"
)

func GetUserById(id string) *database.User {
	return database.GetUser(id)
}

func GetAllUsers() []database.User {
	return database.GetAllUsers()
}

func SaveUser(u database.User) (*database.User, *AppError) {
	if e := validate(u); e != nil {
		return nil, e
	}
	//TODO validar regras de nomes únicos no BD

	if database.GetUser(u.Id) != nil {
		updateUser(u)
	} else {
		u = insertUser(u)
	}

	return &u, nil
}

func updateUser(u database.User) {
	database.PutUser(u)
}

func insertUser(u database.User) database.User {
	if u.Id == "" {
		u.Id = "user-" + strconv.Itoa(1000+rand.Intn(1000))
	}
	database.PutUser(u)
	util.AppLogger.Info("New user created",
		"user", u,
		"totalOfUsers", database.GetTotalUsers())
	return u
}
