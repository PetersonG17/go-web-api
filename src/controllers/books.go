package controllers

import (
	"encoding/json"
	"go-web-api/models"
	"net/http"
)

func ListBooks(response http.ResponseWriter, request *http.Request) {

	// TODO: Fetch from data store
	data := []models.Book{
		{Id: 1, Title: "Some Book", Author: "Terry Nelson"},
		{Id: 2, Title: "Some Book 2", Author: "Pam Anderson"},
		{Id: 3, Title: "Some Book 3", Author: "Loraine Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func GetBook(response http.ResponseWriter, request *http.Request) {
	// TODO: Fetch from data store
	data := []models.Book{
		{Id: 1, Title: "Some Book", Author: "Terry Nelson"},
		{Id: 2, Title: "Some Book 2", Author: "Pam Anderson"},
		{Id: 3, Title: "Some Book 3", Author: "Loraine Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}
