package core

import (
	adapter "allocine-letterboxd-sync-reviews/internal/adapters/allocine"
	entites "allocine-letterboxd-sync-reviews/internal/core/entities"
	"allocine-letterboxd-sync-reviews/internal/core/utils"
	ports "allocine-letterboxd-sync-reviews/internal/ports"
)

type ReviewUseCase struct {
	repository   ports.AllocineRepository
	csvGenerator utils.CsvGenerator
}

func NewReviewUseCase() *ReviewUseCase {
	return &ReviewUseCase{
		repository:   adapter.NewAllocineRepositoryImpl(),
		csvGenerator: *utils.NewCsvGenerator(),
	}
}

func (u *ReviewUseCase) GetAllReviews() []entites.Review {
	return u.repository.GetAllReviews()
}

func (u *ReviewUseCase) SaveReviewsIntoCsv(reviews []entites.Review) {
	u.csvGenerator.Generate(reviews)
}
