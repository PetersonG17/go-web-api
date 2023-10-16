package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-web-api/models"
	"go-web-api/repositories"
	"log"
	"net/http"
	"strconv"
)

type AuthorController struct {
	AuthorRepository repositories.AuthorRepository
}

func ListAuthors(response http.ResponseWriter, request *http.Request) {

	// TODO: Fetch from data store
	data := []models.Author{
		{Id: 1, FirstName: "Terry", LastName: "Nelson"},
		{Id: 2, FirstName: "Pam", LastName: "Anderson"},
		{Id: 3, FirstName: "Lorainer", LastName: "Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func (controller *AuthorController) GetAuthor(response http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		// TODO: log this
		log.Fatal("Can't convert this to an int!")

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)

		data := make(map[string]string)

		data["message"] = "Unable to convert ID to an int"

		json.NewEncoder(response).Encode(data)
	}

	// TODO: Fetch from data store
	author, err := controller.AuthorRepository.Find(parsedId)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(author)
}

func DeleteAuthor(response http.ResponseWriter, request *http.Request) {

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
	data := make(map[string]string)

	data["message"] = "Author with ID: 1 deleted successfully"

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func UpdateAuthor(response http.ResponseWriter, request *http.Request) {

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
	data := models.Author{Id: 1, FirstName: "Terry", LastName: "Clark"}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func CreateAuthor(response http.ResponseWriter, request *http.Request) {

	// TODO: Persist data in the data store
	data := models.Author{Id: 1, FirstName: "Terry", LastName: "Clark"}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}
