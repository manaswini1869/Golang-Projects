# GO Book Management

The Book Management System is a simple RESTful API built in Go for managing books. It allows users to perform CRUD operations (Create, Read, Update, Delete) on book resources through HTTP requests. The project is structured into several packages:

**config**: Contains database configuration and initialization code using the GORM library to interact with a MySQL database.
**controllers**: Defines HTTP request handlers for different CRUD operations on books. These handlers interact with the database through the models package.
**models**: Contains the definition of the Book struct and functions for database operations like creating, reading, updating, and deleting books.
**routes**: Registers HTTP routes using the Gorilla Mux router, mapping them to the corresponding controller functions.
**utils**: Provides utility functions, such as ParseBody, for parsing JSON request bodies.
The main package initializes the HTTP server, sets up routes, and starts listening for incoming requests on a specified port.

Overall, the project provides a foundational structure for building a scalable and maintainable book management system API in Go.
