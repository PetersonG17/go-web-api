package controllers

import (
	"encoding/json"
	"go-web-api/models"
	"net/http"
)

func ListBooks(response http.ResponseWriter, request *http.Request) {

	// TODO: Fetch from data store
	data := []models.Book{
		{Id: "a1e4f70f-883f-462d-ab50-e0c8ac1d9c35", Title: "Some Book", Author: "Terry Nelson"},
		{Id: "c1923cca-db21-48a9-8fdf-f157429481b6", Title: "Some Book 2", Author: "Pam Anderson"},
		{Id: "aa1eeb55-5d65-486e-a698-6f202927fda6", Title: "Some Book 3", Author: "Loraine Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}

func GetBook(response http.ResponseWriter, request *http.Request) {
	// TODO: Fetch from data store
	data := []models.Book{
		{Id: "a1e4f70f-883f-462d-ab50-e0c8ac1d9c35", Title: "Some Book", Author: "Terry Nelson"},
		{Id: "c1923cca-db21-48a9-8fdf-f157429481b6", Title: "Some Book 2", Author: "Pam Anderson"},
		{Id: "aa1eeb55-5d65-486e-a698-6f202927fda6", Title: "Some Book 3", Author: "Loraine Duller"},
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(data)
}
