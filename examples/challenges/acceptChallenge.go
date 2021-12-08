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

	_, err := client.AcceptChallenge("1yTrtw8a")

	if err != nil {
		log.Fatalf("Error accepting challenge: %s", err.Error())
	}

}
