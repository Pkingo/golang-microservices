package app

import (
	"net/http"

	"github.com/pkingo/golang-microservices/mvc/controllers"
)

// StartApp starts the application and listens to port 8080
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
