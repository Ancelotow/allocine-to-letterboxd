package allocine

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

type AllocineGraphQlService struct{}

func NewAllocineGraphQlService() *AllocineGraphQlService {
	return &AllocineGraphQlService{}
}

func (s *AllocineGraphQlService) FetchReview(afterCursor *string, limit int) ResponseData {
	req := graphql.NewRequest(`
        query GetMoviesReviews(
			$after: String
			$first: Int
			$order: UserSocialActionSorting
			$rating: ReviewRating
		) {
			me {
				user {
					id
					social {
						relatedEntities {
							movies(
								entityType: [MOVIE]
								first: $first
								after: $after
								order: [$order]
								rating: $rating
							) {
								totalCount
								pageInfo {
									endCursor
									hasNextPage
								}
								edges {
									node {
										id
										opinion {
											id
											createdAt
											content {
												review
												rating(base: 5)
											}
										}
										movie {
											internalId
											title
											originalTitle
											data {
												productionYear
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
    `)

	req.Var("after", afterCursor)
	req.Var("first", limit)
	req.Var("order", "LATEST")
	req.Var("rating", nil)

	secretKey := os.Getenv("JWT_TOKEN")
	bearerToken := fmt.Sprintf("Bearer %s", secretKey)
	req.Header.Set("Authorization", bearerToken)

	var respData ResponseData
	client := graphql.NewClient("https://graph.allocine.fr/v1/public")
	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData
}
