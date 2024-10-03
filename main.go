package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"Simple-Web-Backend-With-DB-ChatGPT/server"

	"Simple-Web-Backend-With-DB-ChatGPT/models"
)

func main() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("my-database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the database schema
	err = db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router and server
	r := server.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
