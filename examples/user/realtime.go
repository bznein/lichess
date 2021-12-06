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

	realTime, err := client.GetRealTimeUsersStatus([]string{"bznein", "natopigro"}, true)
	if err != nil {
		log.Fatalf("Error getting real time users: %s", err.Error())
	}

	for _, user := range realTime {
		fmt.Printf("Info about %s\n", user.Name)

		if user.Online != nil && *user.Online {
			fmt.Println("Is not online :(")
		} else {
			fmt.Println("Is online")
		}
	}

}
