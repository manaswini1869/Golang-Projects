package config

import (
	_ "githu.com/jinzhu/gorm/dialects/mysql" // Importing MySQL dialect for GORM, but not using it directly in code
	"github.com/jinzhu/gorm"                 // Importing GORM for database management
)

var (
	db *gorm.DB // Global variable to hold the database connection
)

// Connect establishes a connection to the MySQL database
func Connect() {
	// Establishing a connection to the MySQL database
	d, err := gorm.Open("mysql", "manaswini:Manaswini@123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // Panicking if there is an error while establishing the connection
	}
	db = d // Assigning the database connection to the global variable
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db // Returning the database connection
}
