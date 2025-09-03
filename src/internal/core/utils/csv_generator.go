package utils

import (
	"fmt"
	"log"
	"os"
	filepath2 "path/filepath"

	entites "allocine-letterboxd-sync-reviews/internal/core/entities"
)

type CsvGenerator struct{}

func NewCsvGenerator() *CsvGenerator {
	return &CsvGenerator{}
}

func (g *CsvGenerator) Generate(reviews []entites.Review) {
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

	_, err = file.WriteString("AllocineId;MovieTitle;MovieYear;ReviewAt;Rating;Review\n")
	if err != nil {
		log.Fatal("Error writing header to csv file", err)
	}

	for _, review := range reviews {
		currentLine := fmt.Sprintf(
			"%d;\"%s\";%d;%s;%.1f;\"%s\"\n",
			review.AlloCineId,
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
