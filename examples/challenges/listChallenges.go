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

	challenges, err := client.GetChallenges()
	if err != nil {
		log.Fatalf("Error getting challenges: %s", err.Error())
	}

	for _, c := range challenges.Out {
		fmt.Printf("Challenge from %s to %s\n", c.Challenger.Name, c.DestUser.Name)
	}

}
