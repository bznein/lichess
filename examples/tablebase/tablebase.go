package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bznein/lichess"
)

func main() {
	baseUrl, _ := url.Parse("https://tablebase.lichess.ovh")
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
		BaseURL:    baseUrl,
	}

	analysis, err := client.AnalyzeTablebase("standard", "2k5/8/2q5/8/4P3/3P4/2N1N3/2K5 w - - 0 1")
	if err != nil {
		log.Fatalf("Error getting tablebase from  position: %s", err.Error())
	}

	fmt.Printf("This position is a: %s\n", analysis.Category)
}
