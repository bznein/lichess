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

	email, err := client.GetEmail()
	if err != nil {
		log.Fatalf("Error getting email: %s", err.Error())
	}

	fmt.Printf("My account's email is %s\n", email.Email)
}
