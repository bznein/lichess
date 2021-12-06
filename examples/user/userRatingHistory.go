package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bznein/lichess"
)

// NOTE: This is not working currently (see issue #1)
func main() {
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	history, err := client.GetUserRatingHistory("bznein")
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

	fmt.Printf("%v\n", history[0])
}
