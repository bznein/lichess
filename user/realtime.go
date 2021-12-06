package user

type RealTimeUser struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	PlayingId string  `json:"playingId"`
	Title     *string `json:"title"`
	Online    *bool   `json:"online"`
	Playing   *bool   `json:"playing"`
	Streaming *bool   `json:"streaming"`
	Patron    *bool   `json:"patron"`
}

type RealTimeUsers []RealTimeUser
