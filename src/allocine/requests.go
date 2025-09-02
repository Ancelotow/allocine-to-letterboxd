package allocine

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

func GetReviews() []Node {
	req := graphql.NewRequest(`
        query GetMoviesReviews($after: String, $first: Int, $order: UserSocialActionSorting, $rating: ReviewRating, $statuses: [ReviewStatus]!) {
            me {
                user {
                    id
                    social {
                        relatedEntities {
                            moviesReviews: reviews(
                                entityType: [MOVIE]
                                first: $first
                                after: $after
                                order: [$order]
                                rating: $rating
                                statuses: $statuses
                            ) {
                                totalCount
                                pageInfo {
                                    endCursor
                                    hasNextPage
                                }
                                edges {
                                    node {
                                        id
                                        entity {
                                            ... on Movie {
                                                id
                                                internalId
                                                title
                                                data {
                                                    productionYear
                                                }
                                            }
                                        }
                                        opinion {
                                            id
                                            internalId
											createdAt
                                            content {
                                                rating(base: 5)
                                                updatedAt
                                                status
                                                review
                                                helpfulCount
                                                unhelpfulCount
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

	req.Var("after", nil)
	req.Var("first", 50)
	req.Var("order", "LATEST")
	req.Var("rating", nil)
	req.Var("statuses", []string{"PUBLISHED", "ACCEPTED"})

	secretKey := os.Getenv("JWT_TOKEN")
	bearerToken := fmt.Sprintf("Bearer %s", secretKey)
	req.Header.Set("Authorization", bearerToken)

	var respData ResponseData

	client := graphql.NewClient("https://graph.allocine.fr/v1/public")
	if err := client.Run(context.Background(), req, &respData); err != nil {
		log.Fatal(err)
	}

	opinions := respData.Me.User.Social.RelatedEntities.MoviesReviews.Edges
	return opinions
}
