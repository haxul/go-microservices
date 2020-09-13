package dao

import (
	"errors"
	"github.com/haxul/go-microservices/mvc/entities"
)

var users = map[uint64]*entities.User{
	100: {
		Id:        100,
		FirstName: "sergei",
		LastName:  "starodubov",
		Email:     "test@test.com",
	},
}

func GetUser(id uint64) (*entities.User, error) {
	user := users[id]
	if user == nil {
		return &entities.User{}, errors.New("user is not found")
	}
	return user, nil
}
