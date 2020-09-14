package services

import (
	"github.com/haxul/go-microservices/users/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userDaoMock usersDaoMock

type usersDaoMock struct{}

func (u *usersDaoMock) GetUser(id uint64) (*entities.User, error) {
	return &entities.User{Id: id, FirstName: "Test", LastName: "Test", Email: "TEst"}, nil
}

func TestUserService_GetUser(t *testing.T) {
	user, _ := userDaoMock.GetUser(100)
	assert.Equal(t, uint64(100), user.Id)
}
