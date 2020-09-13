package controllers

import (
	"encoding/json"
	"github.com/haxul/go-microservices/mvc/errors"
	"github.com/haxul/go-microservices/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("id")
	id, containsError := strconv.ParseInt(userId, 10, 64)

	if containsError != nil {
		errors.NewError(http.StatusNotFound, containsError, &responseWriter)
		return
	}

	user, err := services.GetUser(uint64(id))
	if err != nil {
		errors.NewError(http.StatusNotFound, err, &responseWriter)
		return
	}

	jsonValue, _ := json.Marshal(user)
	responseWriter.Write(jsonValue)
}
