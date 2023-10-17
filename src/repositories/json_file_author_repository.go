package repositories

import (
	"encoding/json"
	"fmt"
	"go-web-api/helpers"
	"go-web-api/models"
	"log"
	"os"
	"path/filepath"
)

type JsonFileAuthorRepository struct {
}

func (repository JsonFileAuthorRepository) All() ([]models.Author, error) {
	absoluteFilePath, _ := filepath.Abs("./data/authors.json")
	data, err := os.ReadFile(absoluteFilePath)

	if err != nil {
		log.Print(err)
		return []models.Author{}, err
	}

	var authors []models.Author
	json.Unmarshal(data, &authors)

	return authors, nil
}

func (repository JsonFileAuthorRepository) Find(id string) (models.Author, error) {
	absoluteFilePath, _ := filepath.Abs("./data/authors.json")
	file, err := os.Open(absoluteFilePath)

	if err != nil {
		log.Print(err)
		return models.Author{}, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Token()

	var author models.Author

	for decoder.More() {

		decoder.Decode(&author)

		if author.Id == id {
			return author, nil
		}
	}

	// We could not find the author
	return models.Author{}, fmt.Errorf("author with id: %s was not found", id)
}

func (repository JsonFileAuthorRepository) Delete(id string) error {
	absoluteFilePath, _ := filepath.Abs("./data/authors.json")
	data, err := os.ReadFile(absoluteFilePath)

	if err != nil {
		log.Print(err)
		return err
	}

	var authors []models.Author
	json.Unmarshal(data, &authors)

	for i, author := range authors {
		if author.Id == id {
			authors = helpers.Delete(authors, i)
			break
		}
	}

	jsonData, err := json.Marshal(authors)
	if err != nil {
		log.Print(err)
		return err
	}

	os.WriteFile(absoluteFilePath, jsonData, os.ModePerm)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (repository JsonFileAuthorRepository) Save(author models.Author) error {
	absoluteFilePath, _ := filepath.Abs("./data/authors.json")
	data, err := os.ReadFile(absoluteFilePath)

	if err != nil {
		log.Print(err)
		return err
	}

	// Update if exists or save if it does not
	var authors []models.Author
	json.Unmarshal(data, &authors)

	var found bool
	for i, existingAuthor := range authors {

		if existingAuthor.Id == author.Id {
			// TODO: Find a better way to do this, need to make sure that only fields that have been defined can go here
			// Update the author
			if author.FirstName != "" {
				existingAuthor.FirstName = author.FirstName
			}

			if author.LastName != "" {
				existingAuthor.LastName = author.LastName
			}

			authors[i] = existingAuthor
			found = true
			break
		}
	}

	// If it does not exist
	if !found {
		authors = append(authors, author)
	}

	jsonData, err := json.Marshal(authors)
	if err != nil {
		log.Print(err)
		return err
	}

	os.WriteFile(absoluteFilePath, jsonData, os.ModePerm)

	return nil
}
