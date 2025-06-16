package service

import (
	"onboarding/internal/model"
	"testing"
)

// Mock implementations
type mockRepo struct {
	getUserFunc     func(string) *model.User
	getAllUsersFunc func() []model.User
	putUserFunc     func(*model.User) error
}

func (m *mockRepo) GetUser(id string) *model.User {
	return m.getUserFunc(id)
}
func (m *mockRepo) GetAllUsers() []model.User {
	return m.getAllUsersFunc()
}
func (m *mockRepo) PutUser(u *model.User) error {
	return m.putUserFunc(u)
}

type mockValidator struct {
	validateFunc func(*model.User) error
}

func (m *mockValidator) Validate(u *model.User) error {
	return m.validateFunc(u)
}

func TestGetUserById(t *testing.T) {
	//t.Run("User exists", func(t *testing.T) {
	//	mockRepo
	//}
	//repo := &mockRepo{
	//	getUserFunc: func(id string) *model.User {
	//		if id == "1" {
	//			return &model.User{Id: "1", FirstName: "Test"}
	//		}
	//		return nil
	//	},
	//}
	//us := &UserService{repo: &repo, validator: &mockValidator{}}
	//user := us.GetUserById("1")
	//if user == nil || user.Id != "1" {
	//	t.Errorf("Expected user with Id '1', got %v", user)
	//}
}

//func TestGetAllUsers(t *testing.T) {
//	repo := &mockRepo{
//		getAllUsersFunc: func() []model.User {
//			return []model.User{{Id: "a"}, {Id: "b"}}
//		},
//	}
//	us := &UserService{repo: &repo, validator: &mockValidator{}}
//	users := us.GetAllUsers()
//	if len(users) != 2 {
//		t.Errorf("Expected 2 users, got %d", len(users))
//	}
//}
//
//func TestSaveUser_ValidationError(t *testing.T) {
//	repo := &mockRepo{}
//	validator := &mockValidator{
//		validateFunc: func(u *model.User) error { return errors.New("validation error") },
//	}
//	us := &UserService{repo: &repo, validator: validator}
//	err := us.SaveUser(&model.User{})
//	if err == nil || err.Error() != "validation error" {
//		t.Errorf("Expected validation error, got %v", err)
//	}
//}
//
//func TestSaveUser_InsertOrUpdateSuccess(t *testing.T) {
//	called := false
//	repo := &mockRepo{
//		putUserFunc: func(u *model.User) error {
//			called = true
//			return nil
//		},
//	}
//	validator := &mockValidator{
//		validateFunc: func(u *model.User) error { return nil },
//	}
//	us := &UserService{repo: &repo, validator: validator}
//	user := &model.User{}
//	err := us.SaveUser(user)
//	if err != nil {
//		t.Errorf("Expected no error, got %v", err)
//	}
//	if !called {
//		t.Error("Expected PutUser to be called")
//	}
//	if user.Id == "" {
//		t.Error("Expected user Id to be set")
//	}
//}
//
//func TestSaveUser_PutUserError(t *testing.T) {
//	repo := &mockRepo{
//		putUserFunc: func(u *model.User) error { return errors.New("db error") },
//	}
//	validator := &mockValidator{
//		validateFunc: func(u *model.User) error { return nil },
//	}
//	us := &UserService{repo: &repo, validator: validator}
//	user := &model.User{}
//	err := us.SaveUser(user)
//	if err == nil || err.Error() != "db error" {
//		t.Errorf("Expected db error, got %v", err)
//	}
//}
