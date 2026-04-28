package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"url-shortener/database"
	"url-shortener/handlers"
)

func main() {
	var db database.Store
	var err error

	dbType := os.Getenv("DB_TYPE")
	dbURI := os.Getenv("DB_URI")

	switch dbType {
	case "postgres":
		if dbURI == "" {
			dbURI = "postgres://user:password@localhost:5432/urlshortener?sslmode=disable"
		}
		db, err = database.NewPostgresStore(dbURI)
		fmt.Println("Using PostgreSQL database")
	case "mongo":
		if dbURI == "" {
			dbURI = "mongodb://localhost:27017"
		}
		db, err = database.NewMongoStore(dbURI)
		fmt.Println("Using MongoDB database")
	default:
		db = database.NewMemoryStore()
		fmt.Println("Using In-Memory database (default)")
	}

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Routes
	http.HandleFunc("/shorten", handlers.ShortenHandler(db))
	http.HandleFunc("/", handlers.RedirectHandler(db))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("URL Shortener service started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
