package repositories

import "go-web-api/models"

type AuthorRepository interface {
	Find(string) (models.Author, error)
	Save(models.Author) error
	// Get() []models.Author
}
