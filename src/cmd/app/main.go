package main

import (
	"allocine-letterboxd-sync-reviews/internal/core"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	useCase := core.NewReviewUseCase()

	log.Println("Fetching AlloCiné reviews...")
	reviews := useCase.GetAllReviews()
	log.Printf("Found %d AlloCiné reviews\n", len(reviews))

	log.Println("Writing reviews into csv file to import into Letterboxd...")
	useCase.SaveReviewsIntoCsv(reviews)
	log.Println("Reviews written into csv file")
}
