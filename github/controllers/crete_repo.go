package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/haxul/go-microservices/github/model"
	"github.com/haxul/go-microservices/github/services"
)

func CreateRepo(c *gin.Context) {
	repoName := c.Param("name")
	resp, err := services.GithubService.CreateRepo(repoName, "new "+repoName)

	if err != nil {
		c.JSON(400, model.GithubErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, resp)
}
