# GO Book Management

The Book Management System is a simple RESTful API built in Go for managing books. It allows users to perform CRUD operations (Create, Read, Update, Delete) on book resources through HTTP requests. The project is structured into several packages:  

**config**: Contains database configuration and initialization code using the GORM library to interact with a MySQL database.  
**controllers**: Defines HTTP request handlers for different CRUD operations on books. These handlers interact with the database through the models package.  
**models**: Contains the definition of the Book struct and functions for database operations like creating, reading, updating, and deleting books.  
**routes**: Registers HTTP routes using the Gorilla Mux router, mapping them to the corresponding controller functions.  
**utils**: Provides utility functions, such as ParseBody, for parsing JSON request bodies.
The main package initializes the HTTP server, sets up routes, and starts listening for incoming requests on a specified port.  

Overall, the project provides a foundational structure for building a scalable and maintainable book management system API in Go.

# Movies RestFul

This Go project is a simple API for managing movies:

**Data Model:**  

The project defines a Movies struct representing movie data, including fields like ID, Isbn, Title, and Director.  
It also defines a Director struct for storing director information.
API Endpoints:  

GET /movies: Retrieves all movies.  
GET /movies/{id}: Retrieves a specific movie by its ID.    
POST /movies: Creates a new movie.  
PUT /movies/{id}: Updates an existing movie.  
DELETE /movies/{id}: Deletes a movie by its ID.  
Handlers:  

The project defines HTTP request handlers for each endpoint, including getAllMovies, getMovie, createMovie, updateMovie, and deleteMovie.  
These handlers interact with the movie data stored in memory and perform CRUD operations accordingly.  
Router Setup:  

It uses the Gorilla Mux router to define and handle HTTP routes.  
Routes are registered for each API endpoint with corresponding handler functions.  
Data Initialization:  

Initial movie data is provided in the main function for demonstration purposes.  
Server Setup:  

The main function sets up and starts an HTTP server listening on port 8000.  
This project demonstrates a basic implementation of a RESTful API in Go for managing movie  
 data. It can serve as a starting point for building more complex movie management systems.  
