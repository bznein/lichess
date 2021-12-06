package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bznein/lichess"
	"github.com/bznein/lichess/analysis"
)

func main() {
	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	request := analysis.Request{
		Fen:     "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2",
		MultiPV: 5,
		Variant: "standard",
	}

	analysis, err := client.AnalyzePosition(request)
	if err != nil {
		log.Fatalf("Error analyzing position: %s", err.Error())
	}
	fmt.Println(*analysis)

	fmt.Printf("Cached Analysis at %d knodes, depth %d:\n", analysis.Knodes, analysis.Depth)
	fmt.Printf("Best continuation: %s, eval: %f\n", analysis.Pvs[0].Moves, float64(analysis.Pvs[0].Cp)/100.0)
}
