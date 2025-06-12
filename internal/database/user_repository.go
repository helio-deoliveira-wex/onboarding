package database

import (
	"errors"
	"sync"
	"time"
)

// DB mock
var (
	users = make(map[string]User)
	mu    sync.RWMutex
)

func GetUser(id string) *User {
	mu.RLock()
	defer mu.RUnlock()
	u, ok := users[id]
	if !ok {
		return nil
	}
	return &u
}

func GetAllUsers() []User {
	mu.RLock()
	defer mu.RUnlock()
	list := make([]User, 0, len(users))
	for _, v := range users {
		list = append(list, v)
	}
	return list
}

func GetTotalUsers() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(users)
}

var (
	ErrInvalidUser  error = errors.New("user is invalid")
	ErrUserUIsEmpty error = errors.New("user is empty")
)

func PutUser(u *User) error {
	mu.Lock()
	defer mu.Unlock()
	if u == nil {
		return errors.Join(ErrInvalidUser, ErrUserUIsEmpty)
	}
	u.UpdatedAt = time.Now()
	users[u.Id] = *u
	return nil
}
