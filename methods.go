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
	_, err = c.do(req, &user)
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
