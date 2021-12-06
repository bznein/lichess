package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bznein/lichess"
)

func main() {
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	games, err := client.GetOngoingGames()
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

	fmt.Printf("Games in progress:\n")
	for _, g := range games.NowPlaying {
		fmt.Printf("Game link: https://lichess.org/%s\n", g.FullID)
	}
}
