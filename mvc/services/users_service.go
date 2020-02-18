package services

import (
	"github.com/certifiedcloudarchitect/go-mvc/mvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)

}
