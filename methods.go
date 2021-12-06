package lichess

import (
	"fmt"

	"github.com/bznein/lichess/account"
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
