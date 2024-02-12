package main

import (
	"github.com/joho/godotenv"
	"go_crud/internal/http"
	"go_crud/internal/infrastructure/database"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDB()

	server := http.NewServer(db)
	err = server.Start(":3000")

	if err != nil {
		panic(err)
	}
}
