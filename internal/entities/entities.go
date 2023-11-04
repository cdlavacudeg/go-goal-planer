package entities

import (
	"errors"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id:    1,
		Name:  "John",
		Email: "j@j.com",
	},
	{
		Id:    2,
		Name:  "Jane",
		Email: "jane@j.com",
	},
}

func GetUsers() []User {
	return users
}

func GetUser(id int) (*User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func CreateUser(user User) error {
	user.Id = len(users) + 1
	users = append(users, user)
	return nil
}
