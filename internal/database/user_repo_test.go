package database

import (
	"onboarding/internal/model"
	"testing"
	"time"
)

func TestNewUserRepo(t *testing.T) {
	repo := NewUserRepo()
	if repo == nil || repo.users == nil {
		t.Error("Expected initialized UserRepo")
	}
}

func TestPutUser_NilUser(t *testing.T) {
	repo := NewUserRepo()
	err := repo.PutUser(nil)
	if err == nil {
		t.Error("Expected error for nil user")
	}
}

func TestPutUser_AndGetUser(t *testing.T) {
	repo := NewUserRepo()
	user := &model.User{Id: "u1", FirstName: "Test"}
	err := repo.PutUser(user)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	got := repo.GetUser("u1")
	if got == nil || got.Id != "u1" {
		t.Errorf("Expected to retrieve user with Id 'u1', got %v", got)
	}
	if got.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestGetUser_NotFound(t *testing.T) {
	repo := NewUserRepo()
	got := repo.GetUser("not-exist")
	if got != nil {
		t.Error("Expected nil for non-existent user")
	}
}

func TestGetAllUsers(t *testing.T) {
	repo := NewUserRepo()
	repo.PutUser(&model.User{Id: "a"})
	repo.PutUser(&model.User{Id: "b"})
	users := repo.GetAllUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestGetTotalUsers(t *testing.T) {
	repo := NewUserRepo()
	if repo.GetTotalUsers() != 0 {
		t.Error("Expected 0 users initially")
	}
	repo.PutUser(&model.User{Id: "x"})
	if repo.GetTotalUsers() != 1 {
		t.Error("Expected 1 user after insert")
	}
}

func TestPutUser_UpdatesExistingUser(t *testing.T) {
	repo := NewUserRepo()
	user := &model.User{Id: "id", FirstName: "A"}
	repo.PutUser(user)
	time.Sleep(1 * time.Millisecond)
	user2 := &model.User{Id: "id", FirstName: "B"}
	repo.PutUser(user2)
	got := repo.GetUser("id")
	if got.FirstName != "B" {
		t.Errorf("Expected updated FirstName 'B', got '%s'", got.FirstName)
	}
}
