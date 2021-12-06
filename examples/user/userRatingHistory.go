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

	history, err := client.GetUserRatingHistory("bznein")
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

	firstEntry := history[0]
	firstP := firstEntry.Points[0]
	fmt.Printf("In %s, on %d/%d/%d, bznein had a rating of %d\n", firstEntry.Name, firstP[2], firstP[1], firstP[0], firstP[3])
}
