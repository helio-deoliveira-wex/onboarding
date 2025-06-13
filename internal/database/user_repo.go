package database

import (
	"errors"
	"onboarding/internal/model"
	"sync"
	"time"
)

var (
	ErrInvalidUser error = errors.New("user is invalid")
	ErrUserIsEmpty error = errors.New("user is empty")
)

type UserRepo struct {
	users map[string]model.User
	mu    sync.RWMutex
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[string]model.User),
	}
}

func (ur *UserRepo) GetUser(id string) *model.User {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	u, ok := ur.users[id]
	if !ok {
		return nil
	}
	return &u
}

func (ur *UserRepo) GetAllUsers() []model.User {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	list := make([]model.User, 0, len(ur.users))
	for _, v := range ur.users {
		list = append(list, v)
	}
	return list
}

func (ur *UserRepo) GetTotalUsers() int {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	return len(ur.users)
}

func (ur *UserRepo) PutUser(u *model.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	if u == nil {
		return errors.Join(ErrInvalidUser, ErrUserIsEmpty)
	}
	u.UpdatedAt = time.Now()
	ur.users[u.Id] = *u
	return nil
}
