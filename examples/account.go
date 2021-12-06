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

	account, err := client.GetProfile()
	if err != nil {
		log.Fatalf("Error getting profile: %s", err.Error())
	}

	fmt.Printf("Data for %s\n", account.ID)
	fmt.Printf("Bullet score: %d\n", account.Perfs.Bullet.Rating)
	fmt.Printf("Blitz score: %d\n", account.Perfs.Blitz.Rating)
	fmt.Printf("Rapid score: %d\n", account.Perfs.Rapid.Rating)
	fmt.Printf("Classical score: %d\n", account.Perfs.Classical.Rating)
}
