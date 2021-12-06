package challenges

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
	Id            string
	Url           string
	Status        string
	Challenger    Challenger
	DestUser      Challenger
	Variant       Variant
	Rated         bool
	Speed         string
	TimeControl   TimeControl
	Color         string
	Perf          Perf
	Direction     string
	InitialFine   string
	DeclineReason string
}

type Challenges struct {
	In  []Challenge `json:"in"`
	Out []Challenge `json:"out"`
}
