package main

import (
	"log"
	"net/http"
	"os"

	"github.com/meis1kqt/fastapiprac/internal/handlers"
	"github.com/meis1kqt/fastapiprac/internal/storage"
)

func main() {
	
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://taskuser:taskpass@localhost:5432/taskdb?sslmode=disable"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Printf("Starting server %s", serverPort)


	db , err := storage.Connect(databaseURL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	storage := storage.NewTaskStore(db)

	handlers := handlers.NewHandlers(storage)

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", handlers.GetAllTasks)
	mux.HandleFunc("/tasks/", handlers.GetById)

	http.ListenAndServe(":8080", mux)

}