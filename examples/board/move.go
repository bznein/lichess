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

	_, err := client.MakeBoardMove("a32Ymd0v", "e2e4")
	if err != nil {
		log.Fatalf("Error making move: %s", err.Error())
	}

}
