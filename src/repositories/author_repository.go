package repositories

import "go-web-api/models"

type AuthorRepository interface {
	Find(int) (models.Author, error)
	// Get() []models.Author
}
