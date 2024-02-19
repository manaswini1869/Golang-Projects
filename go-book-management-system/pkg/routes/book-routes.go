package routes

import (
	"go-book-management-system/pkg/controllers" // Importing controllers package to register route handlers

	"github.com/gorilla/mux" // Importing mux package for routing
)

// RegisterBookStoreRoutes registers routes for book management operations
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           // Endpoint for creating a new book
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")              // Endpoint for retrieving all books
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // Endpoint for retrieving a specific book by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // Endpoint for updating a specific book by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // Endpoint for deleting a specific book by ID
}
