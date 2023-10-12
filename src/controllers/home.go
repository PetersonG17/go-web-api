package controllers

import (
	"encoding/json"
	"go-web-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, request *http.Request) {

	data := []models.Book{
		{Id: 1, Title: "Some Book", Author: "Terry Nelson"},
		{Id: 2, Title: "Some Book 2", Author: "Pam Anderson"},
		{Id: 3, Title: "Some Book 3", Author: "Loraine Duller"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
