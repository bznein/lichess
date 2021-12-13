package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/bznein/lichess"
	"github.com/bznein/lichess/examples/extras/trie"
	"github.com/bznein/lichess/games"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/guptarohit/asciigraph"
	"github.com/notnil/chess"
)

type moveHistEntry struct {
	move        string
	occurrences int
}

func main() {

	client := lichess.Client{
		HttpClient: &http.Client{},
		Token:      os.Getenv("LICHESS_TOKEN"),
	}

	gameRequest := games.UserGamesRequest{
		Since:     nil,
		Until:     nil,
		Max:       nil,
		Vs:        "",
		Rated:     nil,
		PerfType:  "bullet",
		Color:     "",
		Analysed:  nil,
		Moves:     true,
		PgnInJSON: false,
		Tags:      false,
		Clocks:    false,
		Evals:     false,
		Opening:   false,
		Ongoing:   false,
		Finished:  nil,
		Players:   "",
		Sort:      "dateAsc",
	}

	userName := ""
	fmt.Print("Enter account name: ")
	fmt.Scanf("%s", &userName)

	since := ""
	fmt.Print("Enter start date in format YYYY-MM-DD (defaults to account creation date): ")
	fmt.Scanf("%s", &since)
	if since != "" {
		converted, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00+00:00", since))
		if err != nil {
			log.Fatalf("can't convert %s to timestamp: %s", since, err)
		}
		convertedInt := int(converted.Unix() * 1000)
		gameRequest.Since = &convertedInt
	}

	until := ""
	fmt.Print("Enter end date in format YYYY-MM-DD (defaults to now): ")
	fmt.Scanf("%s", &until)
	if until != "" {
		converted, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00+00:00", until))
		if err != nil {
			log.Fatalf("can't convert %s to timestamp: %s", until, err)
		}
		convertedInt := int(converted.Unix() * 1000)
		gameRequest.Until = &convertedInt
	}

	max := -1
	fmt.Print("Enter max number of games to explore: ")
	fmt.Scanf("%d", &max)
	if max != -1 {
		gameRequest.Max = &max
	}

	vs := ""
	fmt.Print("Enter an opponent to get only games played against them (defaults to all): ")
	fmt.Scanf("%s", &vs)
	gameRequest.Vs = vs

	games, err := client.ExportGamesOfAUser(gameRequest, userName)
	if err != nil {
		log.Fatalf("Error getting games: %s", err.Error())
	}

	latest := games[0].CreatedAt
	earliest := games[len(games)-1].CreatedAt
	bc := widgets.NewBarChart()

	bc.Data = make([]float64, 3)
	rootTrie := trie.NewTrie()
	moves := make([]map[string]int, 0)
	for _, g := range games {

		amIWhite := g.Players.White.User.Id == userName

		for i, m := range strings.Split(g.Moves, " ") {
			if len(moves) <= i {
				moves = append(moves, map[string]int{})
			}
			moves[i][m]++
		}
		rootTrie.Add(strings.ReplaceAll(g.Moves, " ", ""))

		switch g.Status {
		case "stalemate", "draw":
			bc.Data[2]++
		case "mate", "resign", "timeout", "outoftime", "cheat":
			if g.Winner == "white" {
				if amIWhite {
					bc.Data[0]++
				} else {
					bc.Data[1]++
				}
			} else {
				if !amIWhite {
					bc.Data[0]++
				} else {
					bc.Data[1]++
				}
			}
		}

	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// TODO recompute this and update
	bc.Labels = []string{"Wins", "Losses", "Draws"}
	bc.Title = "Winning performance"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	/* ---------------------------------------- */
	movesBC := widgets.NewBarChart()
	movesBC.Title = "Most common move"
	movesBC.SetRect(30, 5, 150, 25)
	movesBC.BarWidth = 5

	advanceMoveHistogram(movesBC, moves[0])

	p1 := widgets.NewPlot()
	p1.Title = "Rating progress"
	p1.Marker = widgets.MarkerDot
	p1.Data = [][]float64{[]float64{}}
	p1.SetRect(5, 25, 150, 50)
	p1.DotMarkerRune = '+'
	p1.AxesColor = ui.ColorWhite
	p1.LineColors[0] = ui.ColorYellow
	p1.DrawDirection = widgets.DrawLeft
	p1.Min = image.Point{5, 25}
	ratingHistory, _ := client.GetUserRatingHistory(userName)
	for _, r := range ratingHistory {
		// TODO nikolas make sure it matches the requirements
		if r.Name != "Bullet" {
			continue
		}
		for _, p := range r.Points {
			converted, err := time.Parse(time.RFC3339, fmt.Sprintf("%d-%02d-%02dT00:00:00+00:00", p[0], p[1]+1, p[2]))
			if err != nil {
				log.Fatalf("can't convert %s to timestamp: %s", until, err)
			}
			convertedInt := int(converted.Unix() * 1000)
			if convertedInt < earliest {
				continue
			}
			if convertedInt > latest {
				continue
			}
			p1.Data[0] = append(p1.Data[0], float64(p[3]))
		}
	}
	ui.Render(bc)
	ui.Render(movesBC)

	p := widgets.NewParagraph()
	p.Text = asciigraph.Plot(p1.Data[0], asciigraph.Height(20), asciigraph.Caption("Rating Progress"), asciigraph.Width(50))
	p.SetRect(5, 26, 100, 53)
	ui.Render(p)

	game := chess.NewGame()
	p2 := widgets.NewParagraph()
	p2.Text = game.Position().Board().Draw()
	p2.SetRect(100, 26, 134, 40)
	ui.Render(p2)

	p3 := widgets.NewParagraph()
	p3.Text = "Use the arrows to move the selector"
	p3.SetRect(5, 54, 80, 58)
	ui.Render(p3)

	selector := widgets.NewParagraph()
	selector.Text = "^"
	selector.SetRect(33, 25, 34, 26)
	selector.Border = false
	ui.Render(selector)

	turn := 0

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Right>":
			r := selector.GetRect()
			selector.SetRect(r.Min.X+6, r.Min.Y, r.Max.X+6, r.Max.Y)
			redrawEverything(bc, movesBC, p, p2, p3, selector)
		case "<Left>":
			r := selector.GetRect()
			selector.SetRect(r.Min.X-6, r.Min.Y, r.Max.X-6, r.Max.Y)
			redrawEverything(bc, movesBC, p, p2, p3, selector)
		case "<Enter>":
			chosen := movesBC.Labels[(selector.Min.X-33)/6]
			game.MoveStr(chosen)
			p2.Text = game.Position().Board().Draw()
			turn++
			// TODO we need to store
			advanceMoveHistogram(movesBC, moves[turn])
			redrawEverything(bc, movesBC, p, p2, p3, selector)
		}
	}
}

func advanceMoveHistogram(hist *widgets.BarChart, moves map[string]int) {
	// TODO do this once and then use it instead of recreating the slices everytime
	// 	firstMoveHistDataWhite := make([]moveHistEntry, len(moves))

	moveHist := make([]moveHistEntry, len(moves))
	i := 0
	for k, v := range moves {
		moveHist[i] = moveHistEntry{
			move:        k,
			occurrences: v,
		}
		i++
	}

	sort.Slice(moveHist, func(i, j int) bool {
		return moveHist[i].occurrences > moveHist[j].occurrences
	})

	hist.Data = make([]float64, len(moveHist))
	hist.Labels = make([]string, len(moveHist))

	for i, v := range moveHist {
		hist.Labels[i] = v.move
		hist.Data[i] = float64(v.occurrences)
	}

}

func redrawEverything(bc ui.Drawable, movesBC ui.Drawable, p ui.Drawable, p2 ui.Drawable, p3 ui.Drawable, selector ui.Drawable) {
	ui.Clear()

	ui.Render(bc)
	ui.Render(movesBC)
	ui.Render(p)
	ui.Render(p2)
	ui.Render(p3)
	ui.Render(selector)
}
