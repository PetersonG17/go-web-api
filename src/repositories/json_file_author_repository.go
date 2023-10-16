package repositories

import (
	"encoding/json"
	"fmt"
	"go-web-api/models"
	"log"
	"os"
	"path/filepath"
)

type JsonFileAuthorRepository struct {
}

func (repository JsonFileAuthorRepository) Find(id int) (models.Author, error) {
	absoluteFilePath, _ := filepath.Abs("./data/authors.json")
	data, err := os.ReadFile(absoluteFilePath)

	if err != nil {
		log.Fatal(err)
		return models.Author{}, err
	}

	var authors []models.Author
	err = json.Unmarshal(data, &authors)

	if err != nil {
		log.Fatal(err)
		return models.Author{}, err
	}

	for i := range authors {
		if authors[i].Id == id {
			return authors[i], nil
		}
	}

	// We could not find the author
	return models.Author{}, fmt.Errorf("Author with ID: %d was not found", id)
}

// func Get() []models.Author, error {

// }
