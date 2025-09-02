package main

import (
	"log"
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

type Review struct {
	AllocineId int
	MovieTitle string
	MovieYear  int
	ReviewAt   time.Time
	Rating     float32
	Review     string
}
