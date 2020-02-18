package domain

import (
	"fmt"
	"github.com/certifiedcloudarchitect/go-mvc/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		2048: {Id: 2048, FirstName: "Git", LastName: "Hub", Email: "github@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
