package lol

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TFT provides methods to interface with tft resource
type TFT struct {
	sling *sling.Sling
}

// NewTFT returns a new TFT
func NewTFT(sling *sling.Sling) *TFT {
	return &TFT{sling: sling.New().Path("tft/")}
}

// Challenger /tft/league/v1/challenger
func (t *TFT) Challenger() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/challenger").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}
