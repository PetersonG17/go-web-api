package routes

import (
	"fmt"
	"go-web-api/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// Authors
	router.Get("/authors", controllers.ListAuthors)
	router.Get("/authors/{id}", controllers.GetAuthor)

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return router
}
