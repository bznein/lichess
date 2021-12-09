package puzzles

import "github.com/bznein/lichess/challenges"

type Player struct {
	Color  string `json:"color"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type Game struct {
	Clock   string          `json:"clock"`
	Id      string          `json:"id"`
	Perf    challenges.Perf `json:"perf"`
	Pgn     string          `json:"pgn"`
	Players []Player        `json:"players"`
	Rated   bool            `json:"rated"`
}

type Puzzle struct {
	Id         string   `json:"id"`
	InitialPly int      `json:"initialPly"`
	Plays      int      `json:"plays"`
	Rating     int      `json:"rating"`
	Solution   []string `json:"solution"`
	Themes     []string `json:"themes"`
}

type PuzzleResult struct {
	Game   Game   `json:"game"`
	Puzzle Puzzle `json:"puzzle"`
}
