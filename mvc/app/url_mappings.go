package app

import (
	"github.com/certifiedcloudarchitect/go-mvc/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
