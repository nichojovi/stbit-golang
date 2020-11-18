package entity

type Movies struct {
	Search      []Movie `json:"Search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDB struct {
	UserID     int64  `json:"user_id" db:"user_id"`
	SearchWord string `json:"search_word" db:"search_word"`
	Pagination int64  `json:"pagination" db:"pagination"`
}
