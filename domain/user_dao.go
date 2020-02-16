package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		1024: &User{Id: 1024, FirstName: "Git", LastName: "Hub", Email: "github@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v was not found", userId))

}
