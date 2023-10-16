package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-web-api/controllers"
	"go-web-api/repositories"
	"net/http"
)

func Register() chi.Router {

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Test
	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("welcome"))
	})

	// Books
	router.Get("/books", controllers.ListBooks)

	// DI for author controller
	authorController := controllers.AuthorController{AuthorRepository: repositories.JsonFileAuthorRepository{}}

	// Authors
	router.Route("/authors", func(router chi.Router) {
		router.Get("/", controllers.ListAuthors)
		router.Get("/{id}", authorController.GetAuthor)
		router.Delete("/{id)}", controllers.DeleteAuthor)
		router.Patch("/{id}", controllers.UpdateAuthor)
		router.Post("/", controllers.CreateAuthor)
	})

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return router
}
