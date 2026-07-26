package main

import (
	"log"

	"backend/internal/database"
	"backend/internal/server"
)

func main() {

	db, err := database.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(db)

	log.Println("Starting server on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
