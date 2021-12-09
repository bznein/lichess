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

	puzzle, err := client.GetDailyPuzzle()
	if err != nil {
		log.Fatalf("Error getting games: %s", err.Error())
	}

	fmt.Println(puzzle)
}
