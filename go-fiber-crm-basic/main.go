package main

import (
	"fmt" // Importing fmt package for formatted I/O

	"go-fiber-crm-basic/database" // Importing database package for database connection
	"go-fiber-crm-basic/lead"     // Importing lead package for lead-related functionalities

	"github.com/gofiber/fiber" // Importing fiber package for building web applications
	"github.com/jinzhu/gorm"   // Importing gorm package for ORM
)

// setupRoutes configures the routes for the application
func setupRoutes(app *fiber.App) {
	// Setting up routes for lead endpoints
	app.Get("api/v1/lead", lead.GetLeads)      // Endpoint to retrieve all leads
	app.Post("api/v1/lead/:id", lead.NewLead)  // Endpoint to create a new lead
	app.Delete("api/v1/lead", lead.DeleteLead) // Endpoint to delete a lead
	app.Get("api/v1/lead/:id", lead.GetLead)   // Endpoint to retrieve a lead by ID
}

// initDatabase initializes the database connection and performs migrations
func initDatabase() {
	var err error
	database.DBCon, err = gorm.Open("sqlite3", "leads.db") // Establishing a connection to the SQLite database
	if err != nil {
		panic("Failed to connect database: ") // Panicking if failed to connect to the database
	}
	fmt.Println("Connection opened to database") // Printing success message after connecting to the database
	database.DBCon.AutoMigrate(&lead.Lead{})     // Auto migrating Lead model to create the necessary table
	fmt.Println("Database Migrated")             // Printing success message after migrating the database
}

func main() {
	app := fiber.New() // Creating a new instance of the Fiber web framework
	initDatabase()     // Initializing the database connection and performing migrations
	setupRoutes(app)   // Setting up routes for the application
	app.Listen(3000)   // Starting the Fiber application on port 3000

	defer database.DBCon.Close() // Closing the database connection when the main function exits
}
