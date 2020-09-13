package services

import (
	"github.com/haxul/go-microservices/users/dao"
	"github.com/haxul/go-microservices/users/entities"
)

func GetUser(id uint64) (*entities.User, error) {
	return dao.GetUser(id)
}
