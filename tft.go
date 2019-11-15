package lol

import (
	"net/http"
)

// Challenger /tft/league/v1/challenger
func (c *Client) Challenger() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("tft/league/v1/challenger").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}
