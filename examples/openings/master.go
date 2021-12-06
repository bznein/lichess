package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bznein/lichess"
	"github.com/bznein/lichess/openings"
)

func main() {
	baseUrl, _ := url.Parse("https://explorer.lichess.ovh")
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
		BaseURL:    baseUrl,
	}

	r := openings.Request{
		Fen:      "rnbqkb1r/ppp2ppp/3p1n2/8/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 4",
		Play:     "f6e4",
		Since:    2018,
		Until:    0,
		Moves:    0,
		TopGames: 0,
	}
	games, err := client.ExploreOpening(r, "masters")
	if err != nil {
		log.Fatalf("Error exploring opening: %s", err.Error())
	}

	topG := games.TopGames[0]
	fmt.Printf("Since 2019, the top game played with this opening was %s vs %s: it was played in %d and the winner was %s\n", topG.White.Name, topG.Black.Name, topG.Year, topG.Game.White.Name)

}
