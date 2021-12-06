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

	perf, err := client.GetUserPerformance("natopigro", "bullet")
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

	fmt.Printf("In bullet, natopigro is better than %f of players!\n", perf.Percentile)
}
