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

	threeCheckLeaderboard, err := client.GetThreeCheckLeaderboard(200)
	if err != nil {
		log.Fatalf("Error getting top10: %s", err.Error())
	}

	for _, r := range threeCheckLeaderboard {
		fmt.Println(r)
	}

}
