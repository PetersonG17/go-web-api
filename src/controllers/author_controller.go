package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-web-api/models"
	"go-web-api/repositories"
	"log"
	"net/http"
)

type AuthorController struct {
	AuthorRepository repositories.AuthorRepository
}

func (controller *AuthorController) ListAuthors(response http.ResponseWriter, request *http.Request) {

	// TODO: Fetch from data store
	data := []models.Author{
		{Id: "43394ae3-7761-4b70-8e8f-8b3b1e1d3041", FirstName: "Terry", LastName: "Nelson"},
		{Id: "0600c07c-9078-4e01-b2f0-17ef3053ba7a", FirstName: "Pam", LastName: "Anderson"},
		{Id: "cee0298a-adac-44dc-bb7d-c87fe6cfce3a", FirstName: "Lorainer", LastName: "Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func (controller *AuthorController) GetAuthor(response http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")

	// if err != nil {
	// 	log.Print("Can't convert this to an int!")

	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.WriteHeader(http.StatusBadRequest)

	// 	data := make(map[string]string)

	// 	data["message"] = "Unable to convert ID to an int"

	// 	json.NewEncoder(response).Encode(data)

	// 	return
	// }

	author, err := controller.AuthorRepository.Find(id)

	if err != nil {
		log.Print(err)

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)

		data := make(map[string]string)

		data["message"] = "Requested author was not found"

		json.NewEncoder(response).Encode(data)

		return
	}

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
	data := models.Author{Id: "3c9600d1-39e9-4f8e-99fe-d34f8b2fa302", FirstName: "Terry", LastName: "Clark"}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func (controller *AuthorController) CreateAuthor(response http.ResponseWriter, request *http.Request) {

	var author models.Author
	err := json.NewDecoder(request.Body).Decode(&author)
	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)

		data := map[string]string{"message": err.Error()}

		json.NewEncoder(response).Encode(data)
		return
	}

	author.Id = uuid.New().String()
	err = controller.AuthorRepository.Save(author)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)

		data := map[string]string{"message": err.Error()}

		json.NewEncoder(response).Encode(data)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	data := make(map[string]string)
	data["message"] = "Author created successfully"
	data["id"] = author.Id

	json.NewEncoder(response).Encode(data)
}
