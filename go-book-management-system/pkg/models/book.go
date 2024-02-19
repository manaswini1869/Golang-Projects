package models

import (
	"go-book-management-system/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book represents the structure of a book in the database
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"` // Name field represents the name of the book
	Author      string `json:"author"`      // Author field represents the author of the book
	Publication string `json:"publication"` // Publication field represents the publication of the book
}

func init() {
	config.Connect()        // Establishing a connection to the database
	db = config.GetDB()     // Getting the database connection
	db.AutoMigrate(&Book{}) // Auto migrating the Book model to create the necessary table
}

// CreateBook creates a new book record in the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBook retrieves all books from the database
func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookByID retrieves a specific book by its ID from the database
func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a specific book from the database
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
