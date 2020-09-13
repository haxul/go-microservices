package services

import (
	"github.com/haxul/go-microservices/mvc/dao"
	"github.com/haxul/go-microservices/mvc/entities"
)

func GetUser(id uint64) (*entities.User, error) {
	return dao.GetUser(id)
}
