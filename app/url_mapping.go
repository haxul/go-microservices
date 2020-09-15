package app

import (
	controllers2 "github.com/haxul/go-microservices/github/controllers"
	"github.com/haxul/go-microservices/users/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/repos/create/:name", controllers2.CreateRepo)
}
