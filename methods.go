package lichess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bznein/lichess/account"
	"github.com/bznein/lichess/analysis"
	"github.com/bznein/lichess/board"
	"github.com/bznein/lichess/challenges"
	"github.com/bznein/lichess/games"
	"github.com/bznein/lichess/openings"
	"github.com/bznein/lichess/tablebase"
	"github.com/bznein/lichess/user"
)

type Ok struct {
	Ok bool `json:"ok"`
}

func (c *Client) GetProfile() (*account.Account, error) {
	req, err := c.newRequest("GET", "/api/account", nil)

	user := &account.Account{}
	_, err = c.do(req, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (c *Client) GetEmail() (*account.Email, error) {
	req, err := c.newRequest("GET", "/api/account/email", nil)

	user := &account.Email{}
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
	req.Header.Del("Content-Type")
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

	req.URL.RawQuery = q.Encode()
	openings := &openings.Opening{}
	_, err = c.do(req, openings)
	if err != nil {
		return nil, err
	}
	return openings, err
}

func (c *Client) AnalyzePosition(request analysis.Request) (*analysis.Analysis, error) {
	req, err := c.newRequest("GET", "/api/cloud-eval", nil)
	q := req.URL.Query()

	if request.Fen != "" {
		q.Add("fen", request.Fen)
	}
	if request.MultiPV != 0 {
		q.Add("multiPv", strconv.Itoa(request.MultiPV))
	}
	if request.Variant != "" {
		q.Add("variant", request.Variant)
	}
	analysis := &analysis.Analysis{}

	req.URL.RawQuery = q.Encode()

	response, err := c.do(req, analysis)
	if err != nil {
		if response.StatusCode == 404 {
			// The position is not cached, this is not an error
			return nil, fmt.Errorf("position analysis not cached")
		}
		return nil, err
	}
	return analysis, err
}

func (c *Client) AnalyzeTablebase(variant string, fen string) (*tablebase.Tablebase, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/%s", variant), nil)
	q := req.URL.Query()

	q.Add("fen", fen)

	req.URL.RawQuery = q.Encode()

	tbEntry := &tablebase.Tablebase{}
	_, err = c.do(req, tbEntry)
	if err != nil {
		return nil, err
	}
	return tbEntry, err
}

func (c *Client) GetRealTimeUsersStatus(ids []string, withGameIds bool) (user.RealTimeUsers, error) {
	req, err := c.newRequest("GET", "/api/users/status", nil)
	q := req.URL.Query()
	q.Add("ids", strings.Join(ids, ","))
	req.URL.RawQuery = q.Encode()

	realTime := user.RealTimeUsers{}
	_, err = c.do(req, &realTime)
	if err != nil {
		return nil, err
	}
	return realTime, err

}

func (c *Client) GetChallenges() (*challenges.Challenges, error) {
	req, err := c.newRequest("GET", "/api/challenge", nil)
	req.Header.Del("Content-Type")
	challenges := &challenges.Challenges{}
	_, err = c.do(req, challenges)
	if err != nil {
		return nil, err
	}
	return challenges, err

}

func (c *Client) GetGameChat(gameId string) (board.Chat, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/chat", gameId), nil)
	req.Header.Del("Content-Type")
	chat := board.Chat{}
	_, err = c.do(req, &chat)
	if err != nil {
		return nil, err
	}
	return chat, err

}

func (c *Client) SendGameMessage(gameId string, room string, text string) (*Ok, error) {
	messagePayload := map[string]string{
		"room": room,
		"text": text,
	}
	jsonPayload, err := json.Marshal(messagePayload)
	if err != nil {
		return nil, fmt.Errorf("can't marshal message into json payload: %s", err)
	}
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/chat", gameId), bytes.NewBuffer(jsonPayload))
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) MakeBoardMove(gameId string, move string) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/move/%s", gameId, move), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) AbortGame(gameId string) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/abort", gameId), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) ResignGame(gameId string) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/abort", gameId), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) HandleDraw(gameId string, accept bool) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/draw/%t", gameId, accept), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) HandleTakeback(gameId string, accept bool) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/takeback/%t", gameId, accept), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) ClaimVictory(gameId string) (*Ok, error) {
	req, err := c.newRequest("POST", fmt.Sprintf("/api/board/game/%s/claim-victory", gameId), nil)
	ok := &Ok{}
	_, err = c.do(req, ok)
	if err != nil {
		return nil, err
	}
	return ok, nil
}

func (c *Client) GetFollowedUsers() (account.Accounts, error) {
	req, err := c.newRequest("GET", "api/rel/following", nil)
	req.Header.Set("Accept", "application/x-ndjson")
	accounts := account.Accounts{}
	_, err = c.do(req, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil

}
