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

	_, err := client.SendGameMessage("79AOf85Di1AB", "player", "GL")
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

}
