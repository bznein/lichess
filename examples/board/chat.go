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

	chat, err := client.GetGameChat("mCDPN794")
	if err != nil {
		log.Fatalf("Error getting chat: %s", err.Error())
	}

	if len(chat) == 0 {
		fmt.Println("There were no messages in this chat!")
	} else {
		for _, message := range chat {
			fmt.Printf("%s:\t%s\n", message.User, message.Text)
		}
	}
}
