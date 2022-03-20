package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the ORM database object.
var DB *gorm.DB

// ConnectDatabase initializes the database connection.
func ConnectDatabase(dbname string) {
	// gorm.Open() will return *gorm.DB object.
	// By default, gorm does connection pooling to support concurrency.
	// So we don't need to do `defer DB.Close()`
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Item{})

	DB = db
}
