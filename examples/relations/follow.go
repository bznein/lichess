package main

import (
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

	_, err := client.FollowPlayer("bzneinBlind")
	if err != nil {
		log.Fatalf("Error following profile: %s", err.Error())
	}

}
