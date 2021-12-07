package challenges

import (
	"net/url"
	"strconv"
)

type Challenger struct {
	Rating      int     `json:"rating"`
	Provisional bool    `json:"provisional"`
	Online      bool    `json:"online"`
	Name        string  `json:"name"`
	Title       *string `json:"title"`
	Patron      bool    `json:"patron"`
	Id          string  `json:"id"`
}

type Variant struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Short string `json:"short"`
}

type TimeControl struct {
	Type        string  `json:"type"`
	Limit       *string `json:"limit"`
	Increment   *int    `json:"increment"`
	Show        *string `json:"show"`
	DaysPerTurn *int    `json:"daysPerTurn"`
}

type Perf struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
}

type Challenge struct {
	ChallengeCommon
	DestUser Challenger `json:"dest_user"`
}

type ChallengeCommon struct {
	Id            string      `json:"id"`
	Url           string      `json:"url"`
	Status        string      `json:"status"`
	Challenger    Challenger  `json:"challenger"`
	Variant       Variant     `json:"variant"`
	Rated         bool        `json:"rated"`
	Speed         string      `json:"speed"`
	TimeControl   TimeControl `json:"time_control"`
	Color         string      `json:"color"`
	Perf          Perf        `json:"perf"`
	Direction     string      `json:"direction"`
	InitialFine   string      `json:"initial_fine"`
	DeclineReason string      `json:"decline_reason"`
}

type Challenges struct {
	In  []Challenge `json:"in"`
	Out []Challenge `json:"out"`
}

type ChallengeRequest struct {
	Rated           bool
	ClockLimit      *int
	ClockIncrement  *int
	Days            *int
	Color           string
	Variant         string
	Fen             string
	KeepAliveStream bool
	AcceptByToken   string
	Message         string
}

func (c ChallengeRequest) IsRealTime() bool {
	return c.ClockLimit != nil && c.ClockIncrement != nil
}

func (c *ChallengeRequest) initDefaults() {
	if c.Color == "" {
		c.Color = "random"
	}
	if c.Variant == "" {
		c.Variant = "standard"
	}
	if c.Fen == "" {
		c.Fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}
	if c.ClockLimit == nil && c.ClockIncrement == nil && c.Days == nil {
		tenMins, fiveSecs := 10, 5
		c.ClockLimit = &tenMins
		c.ClockIncrement = &fiveSecs
	}
	if c.AcceptByToken != "" && c.Message == "" {
		c.Message = "Your game with {opponent} is ready: {game}."
	}
}

func (c *ChallengeRequest) GetRequestData() url.Values {
	c.initDefaults()
	data := url.Values{}

	data.Set("rated", strconv.FormatBool(c.Rated))

	data.Set("color", c.Color)
	data.Set("variant", c.Variant)
	data.Set("fen", c.Fen)
	data.Set("keepAliveStream", strconv.FormatBool(c.KeepAliveStream))
	if c.AcceptByToken != "" {
		data.Set("acceptByToken", c.AcceptByToken)
		data.Set("message", c.Message)
	}
	if c.IsRealTime() {
		data.Set("clock.limit", strconv.Itoa(*c.ClockLimit))
		data.Set("clock.increment", strconv.Itoa(*c.ClockIncrement))
	} else {
		data.Set("days", strconv.Itoa(*c.Days))
	}
	return data
}

type ChallengeResponse struct {
	Challenge ChallengeRet `json:"challenge"`
}

type ChallengeRet struct {
	ChallengeCommon
	DestUser Challenger `json:"destUser"`
}
