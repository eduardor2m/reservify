package main

import (
	"github.com/joho/godotenv"
	"reservify/internal/adapters/delivery/http"
)

func main() {
	if err := godotenv.Load("./cmd/application/.env"); err != nil {
		panic(err)
	}

	api := http.NewAPI(&http.Options{})
	api.Serve()
}
