package openings

type Player struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

type Game struct {
	Id     string `json:"id"`
	Winner string `json:"winner"`
	White  Player `json:"white"`
	Black  Player `json:"black"`
	Year   int    `json:"year"`
	Month  string `json:"month"`
}

type Move struct {
	Uci           string
	San           string
	AverageRating int
	White         int
	Draws         int
	Black         int
	Game          *Game
}

type TopGame struct {
	Uci string `json:"uci"`
	Game
}

type OpeningData struct {
	Eco  string `json:"eco"`
	Name string `json:"name"`
}

type Opening struct {
	White    int         `json:"white"`
	Black    int         `json:"black"`
	Draws    int         `json:"draws"`
	Moves    []Move      `json:"moves"`
	TopGames []TopGame   `json:"topGames"`
	Opening  OpeningData `json:"opening"`
}

type Request struct {
	Fen      string
	Play     string
	Since    int
	Until    int
	Moves    int
	TopGames int
}
