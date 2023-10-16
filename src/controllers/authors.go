package controllers

import (
	"encoding/json"
	"go-web-api/models"
	"net/http"
	// "fmt"
	// "github.com/go-chi/chi/v5"
	// "strconv"
)

func ListAuthors(response http.ResponseWriter, request *http.Request) {

	// TODO: Fetch from data store
	data := []models.Book{
		{Id: 1, Author: "Terry Nelson"},
		{Id: 2, Author: "Pam Anderson"},
		{Id: 3, Author: "Loraine Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func GetAuthor(response http.ResponseWriter, request *http.Request) {

	// FIXME:
	// id := chi.URLParam(request, "id")

	// id, err := strconv.Atoi(id)
	// if err != nil {
	// 	// TODO: log this
	// 	fmt.Println("Can't convert this to an int!")

	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.WriteHeader(http.StatusBadRequest)

	// 	data := make(map[string]string)

	// 	data["message"] = "Unable to convert ID to an int"

	// 	json.NewEncoder(response).Encode(data)
	// }

	// TODO: Fetch from data store
	data := models.Book{Id: 1, Author: "Terry Nelson"}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}
