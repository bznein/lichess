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

	top10, err := client.GetTop10()
	if err != nil {
		log.Fatalf("Error getting top10: %s", err.Error())
	}

	fmt.Printf("In bullet, the best player is %s, with a rating of %d!\n", top10.Bullet[0].ID, top10.Bullet[0].Perfs.Bullet.Rating)
	fmt.Printf("In blitz, the best player is %s, with a rating of %d!\n", top10.Blitz[0].ID, top10.Blitz[0].Perfs.Blitz.Rating)
	fmt.Printf("In rapid, the best player is %s, with a rating of %d!\n", top10.Rapid[0].ID, top10.Rapid[0].Perfs.Rapid.Rating)
	fmt.Printf("In classical, the best player is %s, with a rating of %d!\n", top10.Classical[0].ID, top10.Classical[0].Perfs.Classical.Rating)
	fmt.Printf("In ultrabullet, the best player is %s, with a rating of %d!\n", top10.UltraBullet[0].ID, top10.UltraBullet[0].Perfs.UltraBullet.Rating)
	fmt.Printf("In Chess960, the best player is %s, with a rating of %d!\n", top10.Chess960[0].ID, top10.Chess960[0].Perfs.Chess960.Rating)
	fmt.Printf("In crazyhouse, the best player is %s, with a rating of %d!\n", top10.Crazyhouse[0].ID, top10.Crazyhouse[0].Perfs.Crazyhouse.Rating)
	fmt.Printf("In antichess, the best player is %s, with a rating of %d!\n", top10.Antichess[0].ID, top10.Antichess[0].Perfs.Antichess.Rating)
	fmt.Printf("In atomic, the best player is %s, with a rating of %d!\n", top10.Atomic[0].ID, top10.Atomic[0].Perfs.Atomic.Rating)
	fmt.Printf("In horde, the best player is %s, with a rating of %d!\n", top10.Horde[0].ID, top10.Horde[0].Perfs.Horde.Rating)
	fmt.Printf("In KingOfTheHill, the best player is %s, with a rating of %d!\n", top10.KingOfTheHill[0].ID, top10.KingOfTheHill[0].Perfs.KingOfTheHill.Rating)
	fmt.Printf("In RacingKings, the best player is %s, with a rating of %d!\n", top10.RacingKings[0].ID, top10.RacingKings[0].Perfs.RacingKings.Rating)
	fmt.Printf("In ThreeCheck, the best player is %s, with a rating of %d!\n", top10.ThreeCheck[0].ID, top10.ThreeCheck[0].Perfs.ThreeCheck.Rating)
}
