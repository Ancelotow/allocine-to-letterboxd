package ports

import (
	core "allocine-letterboxd-sync-reviews/internal/core/entities"
)

type AllocineRepository interface {
	GetAllReviews() []core.Review
}
