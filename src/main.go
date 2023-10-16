package main

import (
	"github.com/joho/godotenv"
	"go-web-api/routes"
	"net/http"
	"os"
)

func main() {

	_ = godotenv.Load("../.env")

	router := routes.Register()

	port := os.Getenv("HOST_PORT")
	_ = http.ListenAndServe(":"+port, router)
}
