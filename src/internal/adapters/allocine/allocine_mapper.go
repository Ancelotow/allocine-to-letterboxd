package allocine

import (
	core "allocine-letterboxd-sync-reviews/internal/core/entities"
	"log"
	"time"
)

type AllocineMapper struct{}

func NewAllocineMapper() *AllocineMapper {
	return &AllocineMapper{}
}

func (m *AllocineMapper) Map(opinion Edge) core.Review {
	layout := "2006-01-02T15:04:05"
	createdAt, err := time.Parse(layout, opinion.Node.Opinion.CreatedAt)
	if err != nil {
		log.Fatal("Error parsing date:", err)
	}

	return core.Review{
		AlloCineId: opinion.Node.Movie.InternalID,
		MovieTitle: opinion.Node.Movie.Title,
		MovieYear:  opinion.Node.Movie.Data.ProductionYear,
		ReviewAt:   createdAt,
		Rating:     opinion.Node.Opinion.Content.Rating,
		Review:     opinion.Node.Opinion.Content.Review,
	}
}
