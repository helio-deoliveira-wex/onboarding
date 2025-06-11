package database

// DB mock
var users = make(map[string]User)

func GetUser(id string) *User {
	u, ok := users[id]
	if !ok {
		return nil
	}
	return &u
}

func GetAllUsers() []User {
	list := make([]User, 0, len(users))
	for _, v := range users {
		list = append(list, v)
	}
	return list
}

func GetTotalUsers() int {
	return len(users)
}

func PutUser(u User) {
	users[u.Id] = u
}
