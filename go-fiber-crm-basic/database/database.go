package database

import (
	"github.com/jinzhu/gorm"                   // Importing gorm package for ORM
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Importing SQLite dialect for GORM
)

// DBCon represents the database connection instance
var (
	DBCon *gorm.DB
)
