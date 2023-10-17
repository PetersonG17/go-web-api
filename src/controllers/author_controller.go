package controllers

import (
	"encoding/json"
	"fmt"
	"go-web-api/models"
	"go-web-api/repositories"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AuthorController struct {
	AuthorRepository repositories.AuthorRepository
}

func (controller *AuthorController) ListAuthors(response http.ResponseWriter, request *http.Request) {

	// TODO: Implement Pagination and Filters in Query String
	authors, err := controller.AuthorRepository.All()

	if err != nil {
		log.Print(err)

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)

		data := map[string]string{"message": err.Error()}

		json.NewEncoder(response).Encode(data)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(authors)
}

func (controller *AuthorController) GetAuthor(response http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")

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

func (controller *AuthorController) DeleteAuthor(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	err := controller.AuthorRepository.Delete(id)

	if err != nil {
		log.Print(err)

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)

		data := map[string]string{"message": err.Error()}

		json.NewEncoder(response).Encode(data)

		return
	}

	data := make(map[string]string)

	data["message"] = fmt.Sprintf("Author with ID: %s deleted successfully", id)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func (controller *AuthorController) UpdateAuthor(response http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")

	var author models.Author
	err := json.NewDecoder(request.Body).Decode(&author)
	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)

		data := map[string]string{"message": err.Error()}

		json.NewEncoder(response).Encode(data)
		return
	}

	author.Id = id
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
	data["message"] = "Author updated successfully"
	data["id"] = author.Id

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
	response.WriteHeader(http.StatusCreated)

	data := make(map[string]string)
	data["message"] = "Author created successfully"
	data["id"] = author.Id

	json.NewEncoder(response).Encode(data)
}
