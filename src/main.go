package main

import (
	"log"

	"allocine-letterboxd-sync-reviews/allocine"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	allocine.GetReviews()

}
