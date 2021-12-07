package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bznein/lichess"
	"github.com/bznein/lichess/challenges"
)

func main() {
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	ten := 60
	request := challenges.ChallengeRequest{
		Rated:          false,
		ClockLimit:     &ten,
		ClockIncrement: &ten,
		Days:           nil,
		Color:          "random",
		Variant:        "chess960",
	}
	challenge, err := client.CreateChallenge(request, "bzneinBlind")
	if err != nil {
		log.Fatalf("Error creating challenge: %s", err.Error())
	}

	fmt.Printf("Challenge from %s to %s\n", challenge.Challenger.Name, challenge.DestUser.Name)

}
