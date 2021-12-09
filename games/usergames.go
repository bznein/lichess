package games

import (
	"net/url"
	"strconv"
)

type UserGamesRequest struct {
	Since     *int
	Until     *int
	Max       *int
	Vs        string
	Rated     *bool
	PerfType  string
	Color     string
	Analysed  *bool
	Moves     bool
	PgnInJSON bool
	Tags      bool
	Clocks    bool
	Evals     bool
	Opening   bool
	Ongoing   bool
	Finished  *bool
	Players   string
	Sort      string
}

func (r UserGamesRequest) GetRequestData() url.Values {
	data := url.Values{}

	if r.Since != nil {
		data.Set("since", strconv.FormatInt(int64(*r.Since), 10))
	}

	if r.Until != nil {
		data.Set("until", strconv.FormatInt(int64(*r.Until), 10))
	}

	if r.Max != nil {
		data.Set("max", strconv.FormatInt(int64(*r.Max), 10))
	}

	if r.Vs != "" {
		data.Set("vs", r.Vs)
	}

	if r.Rated != nil {
		data.Set("rated", strconv.FormatBool(*r.Rated))
	}

	if r.PerfType != "" {
		data.Set("perfType", r.PerfType)
	}

	if r.Color != "" {
		data.Set("color", r.Color)
	}

	if r.Analysed != nil {
		data.Set("analysed", strconv.FormatBool(*r.Analysed))
	}

	data.Set("moves", strconv.FormatBool(r.Moves))
	data.Set("pgnInJSON", strconv.FormatBool(r.PgnInJSON))
	data.Set("tags", strconv.FormatBool(r.Tags))
	data.Set("clocks", strconv.FormatBool(r.Clocks))
	data.Set("evals", strconv.FormatBool(r.Evals))
	data.Set("opening", strconv.FormatBool(r.Opening))
	data.Set("ongoing", strconv.FormatBool(r.Ongoing))

	if r.Finished != nil {
		data.Set("finished", strconv.FormatBool(*r.Finished))
	}

	if r.Players != "" {
		data.Set("players", r.Players)
	}
	if r.Sort != "" {
		data.Set("sort", r.Sort)
	}
	return data
}
