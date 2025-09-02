package allocine

type MovieData struct {
	ProductionYear int `json:"productionYear"`
}

type Movie struct {
	ID         string    `json:"id"`
	InternalID int       `json:"internalId"`
	Title      string    `json:"title"`
	Data       MovieData `json:"data"`
}

type ReviewContent struct {
	Rating         float32 `json:"rating"`
	UpdatedAt      string  `json:"updatedAt"`
	Status         string  `json:"status"`
	Review         string  `json:"review"`
	HelpfulCount   int     `json:"helpfulCount"`
	UnhelpfulCount int     `json:"unhelpfulCount"`
}

type Opinion struct {
	ID         string        `json:"id"`
	InternalID int           `json:"internalId"`
	CreatedAt  string        `json:"createdAt"`
	Content    ReviewContent `json:"content"`
}

type ReviewNode struct {
	ID      string  `json:"id"`
	Entity  Movie   `json:"entity"`
	Opinion Opinion `json:"opinion"`
}

type Node struct {
	Node ReviewNode `json:"node"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type Reviews struct {
	TotalCount int      `json:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"`
	Edges      []Node   `json:"edges"`
}

type RelatedEntities struct {
	MoviesReviews Reviews `json:"moviesReviews"`
}

type Social struct {
	RelatedEntities RelatedEntities `json:"relatedEntities"`
}

type User struct {
	ID     string `json:"id"`
	Social Social `json:"social"`
}

type Me struct {
	User User `json:"user"`
}

type ResponseData struct {
	Me Me `json:"me"`
}
