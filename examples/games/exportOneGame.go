package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bznein/lichess"
	"github.com/bznein/lichess/games"
)

func main() {
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	gameRequest := games.GameRequest{
		Moves:     true,
		PgnInJson: true,
		Tags:      true,
		Clocks:    true,
		Evals:     true,
		Opening:   true,
		Literate:  true,
		Players:   "",
	}

	game, err := client.ExportOneGame(gameRequest, "a3JjZYYp")
	if err != nil {
		log.Fatalf("Error getting game: %s", err.Error())
	}

	fmt.Println(game)
}
