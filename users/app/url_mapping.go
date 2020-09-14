package app

import (
	"github.com/haxul/go-microservices/users/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
