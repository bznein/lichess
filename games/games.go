package games

type Variant struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Opponent struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Rating   int    `json:"rating"`
}

type Game struct {
	FullID   string   `json:"fullID"`
	GameID   string   `json:"gameID"`
	Fen      string   `json:"fen"`
	Color    string   `json:"color"`
	LastMove string   `json:"lastMove"`
	Variant  Variant  `json:"variant"`
	Speed    string   `json:"speed"`
	Perf     string   `json:"perf"`
	Rated    bool     `json:"rated"`
	Opponent Opponent `json:"opponent"`
	IsMyTurn bool     `json:"isMyTurn"`
}

type Games struct {
	NowPlaying []Game `json:"nowPlaying"`
}
