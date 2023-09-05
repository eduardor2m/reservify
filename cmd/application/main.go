package main

import (
	"log"
	"reservify/internal/adapters/delivery/http"
	"reservify/internal/adapters/persistence/postgres"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./cmd/application/.env"); err != nil {
		log.Println("Error loading .env file")
	}

	postgres.RunMigrations()

	api := http.NewAPI(&http.Options{})
	api.Serve()
}
