package service

import (
	"math/rand"
	"strconv"

	"onboarding/internal/database"
	"onboarding/internal/util"

	"github.com/pkg/errors"
)

type UserRepositoryInterface interface {
	GetUser(id string) *database.User
	GetAllUsers() []database.User
	GetTotalUsers() int
	PutUser(user *database.User) error
}

type IUserService interface {
	GetUserById(id string) *database.User
	GetAllUsers() []database.User
	SaveUser(user database.User) (*database.User, error)
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetUserById(id string) *database.User {
	return us.repo.GetUser(id)
}

func (us *UserService) GetAllUsers() []database.User {
	return us.repo.GetAllUsers()
}

func (us *UserService) SaveUser(u database.User) (*database.User, error) {
	if e := validate(u); e != nil {
		return nil, errors.Wrap(e, "Error saving user")
	}
	// TODO validar regras de nomes únicos no BD

	if us.repo.GetUser(u.Id) != nil {
		us.updateUser(u)
	} else {
		u = us.insertUser(u)
	}

	return &u, nil
}

func (us *UserService) updateUser(u database.User) {
	us.repo.PutUser(&u)
}

func (us *UserService) insertUser(u database.User) database.User {
	if u.Id == "" {
		u.Id = "user-" + strconv.Itoa(1000+rand.Intn(1000))
	}
	us.repo.PutUser(&u)
	util.AppLogger.Info("New user created",
		"user", u,
		"totalOfUsers", us.repo.GetTotalUsers())
	return u
}
