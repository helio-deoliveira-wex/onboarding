package database

import (
	"errors"
	"sync"
	"time"
)

type UserRepo struct {
	users map[string]User
	mu    sync.RWMutex
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[string]User),
	}
}

func (ur *UserRepo) GetUser(id string) *User {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	u, ok := ur.users[id]
	if !ok {
		return nil
	}
	return &u
}

func (ur *UserRepo) GetAllUsers() []User {
	ur.mu.RLock()
	defer ur.mu.RUnlock()
	list := make([]User, 0, len(ur.users))
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

var (
	ErrInvalidUser  error = errors.New("user is invalid")
	ErrUserUIsEmpty error = errors.New("user is empty")
)

func (ur *UserRepo) PutUser(u *User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	if u == nil {
		return errors.Join(ErrInvalidUser, ErrUserUIsEmpty)
	}
	u.UpdatedAt = time.Now()
	ur.users[u.Id] = *u
	return nil
}
