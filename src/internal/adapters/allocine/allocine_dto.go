package allocine

type MovieData struct {
	ProductionYear int `json:"productionYear"`
}

type Movie struct {
	ID            string    `json:"id"`
	InternalID    int       `json:"internalId"`
	Title         string    `json:"title"`
	OriginalTitle string    `json:"originalTitle"`
	Data          MovieData `json:"data"`
}

type ReviewContent struct {
	Rating float32 `json:"rating"`
	Review string  `json:"review"`
}

type Opinion struct {
	ID        string        `json:"id"`
	CreatedAt string        `json:"createdAt"`
	Content   ReviewContent `json:"content"`
}

type Node struct {
	ID      string  `json:"id"`
	Movie   Movie   `json:"movie"`
	Opinion Opinion `json:"opinion"`
}

type Edge struct {
	Node Node `json:"node"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type Movies struct {
	TotalCount int      `json:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"`
	Edges      []Edge   `json:"edges"`
}

type RelatedEntities struct {
	Movies Movies `json:"movies"`
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
