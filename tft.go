package lol

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TFTService provides methods to interface with tft resource
type TFTService struct {
	sling *sling.Sling
}

// NewTFTService returns a new LOLService
func NewTFTService(sling *sling.Sling) *TFTService {
	return &TFTService{sling: sling.New().Path("tft/")}
}

// Challenger /tft/league/v1/challenger
func (t *TFTService) Challenger() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/challenger").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}
