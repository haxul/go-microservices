package dao

import (
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	_, err := GetUser(0)
	if err == nil {
		t.Error("user must be not found")
	}

	user, errUser := GetUser(100)
	if errUser != nil {
		t.Error("user must by found")
	}

	if user.Id != 100 {
		t.Error("incorrect user id")
	}
}
