package games

import (
	"net/url"
	"strconv"
)

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

type GameRequest struct {
	Moves     bool
	PgnInJson bool
	Tags      bool
	Clocks    bool
	Evals     bool
	Opening   bool
	Literate  bool
	Players   string
}

func (g GameRequest) GetQueryData() url.Values {
	data := url.Values{}
	data.Add("moves", strconv.FormatBool(g.Moves))
	data.Add("pgnInJson", strconv.FormatBool(g.PgnInJson))
	data.Add("tags", strconv.FormatBool(g.Tags))
	data.Add("clocks", strconv.FormatBool(g.Clocks))
	data.Add("evals", strconv.FormatBool(g.Evals))
	data.Add("literate", strconv.FormatBool(g.Literate))
	data.Add("opening", strconv.FormatBool(g.Opening))
	if g.Players != "" {
		data.Add("players", g.Players)
	}
	return data
}

type User struct {
	Name   string  `json:"name"`
	Title  *string `json:"title"`
	Patron bool    `json:"patron"`
	Id     string  `json:"id"`
}

type Analysis struct {
	Inaccuracy int `json:"inaccuracy"`
	Mistake    int `json:"mistake"`
	Blunder    int `json:"blunder"`
	Acpl       int `json:"acpl"`
}

type Player struct {
	User        User     `json:"user"`
	Rating      int      `json:"rating"`
	RatingDiff  int      `json:"ratingDiff"`
	Name        string   `json:"name"`
	Provisional bool     `json:"provisional"`
	AiLevel     int      `json:"aiLevel"`
	Analysis    Analysis `json:"analysis"`
	Team        string   `json:"team"`
}

type Players struct {
	White Player `json:"white"`
	Black Player `json:"black"`
}

type Opening struct {
	Eco  string `json:"eco"`
	Name string `json:"name"`
	Ply  int    `json:"ply"`
}

type Judgment struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Eval struct {
	Eval      int      `json:"eval"`
	Best      string   `json:"best"`
	Variation string   `json:"variation"`
	Judgment  Judgment `json:"judgment"`
}

type Clock struct {
	Initial   int `json:"initial"`
	Increment int `json:"increment"`
	TotalTime int `json:"totalTime"`
}

type GameResult struct {
	CommonGameResult
	InitialFen  string `json:"initialFen"`
	Winner      string `json:"winner"`
	Pgn         string `json:"pgn"`
	DaysPerTurn int    `json:"daysPerTurn"`
	Analysis    []Eval `json:"analysis"`
	Tournament  string `json:"tournament"`
	Swiss       string `json:"swiss"`
}

type CommonGameResult struct {
	Id         string  `json:"id"`
	Rated      bool    `json:"rated"`
	Variant    string  `json:"variant"`
	Speed      string  `json:"speed"`
	Perf       string  `json:"perf"`
	CreatedAt  int     `json:"createdAt"`
	LastMoveAt int     `json:"lastMoveAt"`
	Status     string  `json:"status"`
	Players    Players `json:"players"`
	Opening    Opening `json:"opening"`
	Moves      string  `json:"moves"`
	Clock      Clock   `json:"clock"`
}
