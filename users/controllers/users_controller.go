package controllers

import (
	"encoding/json"
	"errors"
	"github.com/haxul/go-microservices/users/custumErrors"
	"github.com/haxul/go-microservices/users/services"
	"net/http"
	"strconv"
)

func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("id")
	id, containsError := strconv.ParseInt(userId, 10, 64)

	if containsError != nil {
		custumErrors.NewError(http.StatusBadRequest, errors.New("must be a number"), &responseWriter)
		return
	}

	user, err := services.GetUser(uint64(id))
	if err != nil {
		custumErrors.NewError(http.StatusBadRequest, err, &responseWriter)
		return
	}

	jsonValue, _ := json.Marshal(user)
	_, jsonError := responseWriter.Write(jsonValue)

	if jsonError != nil {
		custumErrors.NewError(http.StatusBadRequest, jsonError, &responseWriter)
	}
}