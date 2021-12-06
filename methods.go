package lichess

import (
	"fmt"
	"strconv"

	"github.com/bznein/lichess/account"
	"github.com/bznein/lichess/games"
	"github.com/bznein/lichess/openings"
	"github.com/bznein/lichess/user"
)

func (c *Client) GetProfile() (*account.Account, error) {
	req, err := c.newRequest("GET", "/api/account", nil)

	user := &account.Account{}
	_, err = c.do(req, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *Client) GetUser(username string) (*user.User, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s", username), nil)

	user := &user.User{}
	_, err = c.do(req, user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *Client) GetUserRatingHistory(username string) (user.RatingHistory, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s/rating-history", username), nil)

	user := user.RatingHistory{}
	_, err = c.do(req, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *Client) GetUserPerformance(username string, gameType string) (*user.Performance, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s/perf/%s", username, gameType), nil)

	perf := &user.Performance{}
	_, err = c.do(req, perf)
	if err != nil {
		return nil, err
	}
	return perf, err
}

func (c *Client) GetOngoingGames() (*games.Games, error) {
	req, err := c.newRequest("GET", "/api/account/playing", nil)

	games := &games.Games{}
	_, err = c.do(req, games)
	if err != nil {
		return nil, err
	}
	return games, err
}

func (c *Client) GetTop10() (*user.Top10, error) {
	req, err := c.newRequest("GET", "/player", nil)
	req.Header.Set("Accept", "application/vnd.lichess.v3+json")
	top10 := &user.Top10{}
	_, err = c.do(req, top10)
	if err != nil {
		return nil, err
	}
	return top10, err
}

func (c *Client) ExploreOpening(request openings.Request, gameType string) (*openings.Opening, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s", gameType), nil)
	q := req.URL.Query()

	if request.Fen != "" {
		q.Add("fen", request.Fen)
	}
	if request.Play != "" {
		q.Add("play", request.Play)
	}
	if request.Since != 0 {
		q.Add("since", strconv.Itoa(request.Since))
	}
	if request.Until != 0 {
		q.Add("until", strconv.Itoa(request.Until))
	}
	if request.Moves != 0 {
		q.Add("moves", strconv.Itoa(request.Moves))
	}
	if request.TopGames != 0 {
		q.Add("topGames", strconv.Itoa(request.TopGames))
	}

	openings := &openings.Opening{}
	_, err = c.do(req, openings)
	if err != nil {
		return nil, err
	}
	return openings, err
}
