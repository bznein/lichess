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

	gameRequest := games.UserGamesRequest{}

	games, err := client.ExportGamesOfAUser(gameRequest, "bznein")
	if err != nil {
		log.Fatalf("Error getting games: %s", err.Error())
	}

	fmt.Println(games)
}
