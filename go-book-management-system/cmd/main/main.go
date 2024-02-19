package main

import (
	"go-book-management-system/pkg/routes" // Importing routes package which contains route handlers
	"log"                                  // Importing log package for logging errors
	"net/http"                             // Importing net/http package for HTTP server functionalities

	"github.com/gorilla/mux"                  // Importing mux package for routing
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing MySQL dialect for GORM, but not using it directly in code
)

func main() {
	r := mux.NewRouter()                                // Creating a new router using mux
	routes.RegisterBookStoreRoutes(r)                   // Registering routes defined in the routes package
	http.Handle("/", r)                                 // Handling requests with the router
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // Starting HTTP server on port 9010 and logging fatal errors if any
}
