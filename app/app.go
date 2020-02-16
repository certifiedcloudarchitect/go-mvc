package app

import (
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	http.ListenAndServe(":8080", nil)
}
