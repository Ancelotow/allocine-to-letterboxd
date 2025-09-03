package allocine

import (
	core "allocine-letterboxd-sync-reviews/internal/core/entities"
	"allocine-letterboxd-sync-reviews/internal/ports"
)

type AllocineRepositoryImpl struct {
	service *AllocineGraphQlService
	mapper  *AllocineMapper
}

func NewAllocineRepositoryImpl() ports.AllocineRepository {
	return &AllocineRepositoryImpl{
		service: NewAllocineGraphQlService(),
		mapper:  NewAllocineMapper(),
	}
}

func (r *AllocineRepositoryImpl) GetAllReviews() []core.Review {
	var opinions []Edge
	var currentEndCursor *string = nil
	for {
		respData := r.service.FetchReview(currentEndCursor, 5)
		opinions = append(opinions, respData.Me.User.Social.RelatedEntities.Movies.Edges...)
		if !respData.Me.User.Social.RelatedEntities.Movies.PageInfo.HasNextPage {
			break
		}
		currentEndCursor = &respData.Me.User.Social.RelatedEntities.Movies.PageInfo.EndCursor
	}

	var results []core.Review
	for _, opinion := range opinions {
		results = append(results, r.mapper.Map(opinion))
	}

	return results
}
