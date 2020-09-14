package services

import (
	"github.com/haxul/go-microservices/users/dao"
	"github.com/haxul/go-microservices/users/entities"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(id uint64) (*entities.User, error) {
	return dao.UserDao.GetUser(id)
}
