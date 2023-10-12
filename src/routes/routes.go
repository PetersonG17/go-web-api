package routes

import (
	"go-web-api/controllers"
	"net/http"
)

func Register() {
	http.HandleFunc("/", controllers.Home)
}
