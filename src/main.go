package main

import (
	"fmt"
	"log"
	"os"
	filepath2 "path/filepath"
	"time"

	"allocine-letterboxd-sync-reviews/allocine"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opinions := allocine.GetReviews()
	reviews := make([]Review, len(opinions))
	for i, opinion := range opinions {
		reviews[i] = mapToReview(opinion)
	}

	log.Printf("Fetched %d reviews from Allocine", len(reviews))
	saveIntoCsv(reviews)
}

func mapToReview(opinion allocine.Node) Review {
	layout := "2006-01-02T15:04:05"
	createdAt, err := time.Parse(layout, opinion.Node.Opinion.CreatedAt)
	if err != nil {
		log.Fatal("Error parsing date:", err)
	}

	return Review{
		AllocineId: opinion.Node.Entity.InternalID,
		MovieTitle: opinion.Node.Entity.Title,
		MovieYear:  opinion.Node.Entity.Data.ProductionYear,
		ReviewAt:   createdAt,
		Rating:     opinion.Node.Opinion.Content.Rating,
		Review:     opinion.Node.Opinion.Content.Review,
	}
}

func saveIntoCsv(reviews []Review) {
	outputDir := "output"
	outputFile := "allocine-reviews.csv"
	outputPath := filepath2.Join(outputDir, outputFile)

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal("Error creating output directory", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal("Error creating csv file", err)
	}

	_, err = file.WriteString("AllocineId,MovieTitle,MovieYear,ReviewAt,Rating,Review\n")
	if err != nil {
		log.Fatal("Error writing header to csv file", err)
	}

	currentLine := ""
	for _, review := range reviews {
		currentLine = fmt.Sprintf(
			"%d;%s;%d;%s;%.1f;%s\n",
			review.AllocineId,
			review.MovieTitle,
			review.MovieYear,
			review.ReviewAt.Format("2006-01-02"),
			review.Rating,
			review.Review,
		)
		_, err := file.WriteString(currentLine)
		if err != nil {
			log.Fatal("Error writing to csv file", err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing csv file", err)
		}
	}(file)
}

type Review struct {
	AllocineId int
	MovieTitle string
	MovieYear  int
	ReviewAt   time.Time
	Rating     float32
	Review     string
}
