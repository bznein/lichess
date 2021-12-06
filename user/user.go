package user

import "github.com/bznein/lichess/account"

type User struct {
	account.Account
}

type RatingEntry struct {
	Name   string
	Points [][]int
}

type RatingHistory []RatingHistory
