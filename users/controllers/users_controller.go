package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/haxul/go-microservices/users/custumErrors"
	"github.com/haxul/go-microservices/users/services"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	id, containsError := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if containsError != nil {
		c.JSON(http.StatusBadRequest, custumErrors.AppError{Status: http.StatusBadRequest, Message: "must be a number"})
		return
	}

	user, err := services.UserService.GetUser(uint64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, custumErrors.AppError{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
