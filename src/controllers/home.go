package controllers

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(writer, "Made it here")
}
