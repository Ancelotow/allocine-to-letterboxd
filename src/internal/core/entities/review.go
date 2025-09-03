package entities

import "time"

type Review struct {
	AlloCineId    int
	MovieTitle    string
	OriginalTitle string
	MovieYear     int
	ReviewAt      time.Time
	Rating        float32
	Review        string
}
