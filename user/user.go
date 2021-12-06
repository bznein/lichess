package user

import "github.com/bznein/lichess/account"

type User struct {
	account.Account
}

type RatingEntry struct {
	Name   string
	Points [][]int
}

type RatingHistory []RatingEntry

type Glicko struct {
	Rating      float64 `json:"rating"`
	Deviation   float64 `json:"deviation"`
	Provisional bool    `json:"provisional"`
}

type Perf struct {
	NB       int `json:"nb"`
	Progress int `json:"progress"`
}

type PerfType struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type HighLow struct {
	Int    int    `json:"int"`
	At     string `json:"at"`
	GameId string `json:"gameId"`
}

type OpId struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Title *string `json:"title"`
}

type Result struct {
	OpInt  int    `json:"opInt"`
	OpId   OpId   `json:"opId"`
	At     string `json:"at"`
	GameId string `json:"gameId"`
}

type BestWins struct {
	Results []Result `json:"results"`
}

type WorstLosses struct {
	Results []Result `json:"results"`
}

type Cur struct {
	V int `json:"v"`
}

type StreakEndpoint struct {
	At     string `json:"at"`
	GameId string `json:"gameId"`
}

type Max struct {
	V    int            `json:"v"`
	From StreakEndpoint `json:"from"`
	To   StreakEndpoint `json:"to"`
}

type Streak struct {
	Cur Cur `json:"cur"`
}

type ResultStreak struct {
	Win  Streak `json:"win"`
	Loss Streak `json:"loss"`
}

type Nb struct {
	Cur Cur `json:"cur"`
	Max Max `json:"max"`
}

type Time struct {
	Nb
}

type Count struct {
	All         int     `json:"all"`
	Rated       int     `json:"rated"`
	Win         int     `json:"win"`
	Loss        int     `json:"loss"`
	Draw        int     `json:"draw"`
	Tour        int     `json:"tour"`
	Berserk     int     `json:"berserk"`
	OpAvg       float64 `json:"opAvg"`
	Seconds     int     `json:"seconds"`
	Disconnects int     `json:"disconnects"`
}
type PlayStreak struct {
	Nb       Nb     `json:"nb"`
	Time     Time   `json:"time"`
	LastDate string `json:"lastDate"`
}

type Stat struct {
	PerfType     PerfType     `json:"perfType"`
	Highest      HighLow      `json:"highest"`
	Lowest       HighLow      `json:"lowest"`
	BestWins     BestWins     `json:"bestWins"`
	WorstLosses  WorstLosses  `json:"worstLosses"`
	Count        Count        `json:"count"`
	ResultStreak ResultStreak `json:"resultStreak"`
	PlayStreak   PlayStreak   `json:"playStreak"`
}

type Performance struct {
	Perf       Perf    `json:"perf"`
	Rank       int     `json:"rank"`
	Percentile float64 `json:"percentile"`
	Stat       Stat    `json:"stat"`
}
