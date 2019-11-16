package lol

import (
	"fmt"
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

// Challenger GET /tft/league/v1/challenger
func (t *TFT) Challenger() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/challenger").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// EntriesBySummoner GET /tft/league/v1/entries/by-summoner/{encryptedSummonerID}
func (t *TFT) EntriesBySummoner(encryptedSummonerID string) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/entries/by-summoner/"+encryptedSummonerID).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// Entries GET /tft/league/v1/entries/{tier}/{division}
func (t *TFT) Entries(tier, division string, params *EntriesParams) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	endpoint := fmt.Sprintf("league/v1/entries/%s/%s", tier, division)
	resp, err := t.sling.Get(endpoint).QueryStruct(params).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// Grandmaster GET /tft/league/v1/grandmaster
func (t *TFT) Grandmaster() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/grandmaster").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Leagues GET /tft/league/v1/leagues/{leagueID}
func (t *TFT) Leagues(leagueID string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/leagues/"+leagueID).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Master GET /tft/league/v1/master
func (t *TFT) Master() (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := t.sling.Get("league/v1/master").Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// MatchesByPUUID GET /tft/match/v1/matches/by-puuid/{encryptedPUUID}/ids
func (t *TFT) MatchesByPUUID(encryptedPUUID string) ([]string, *http.Response, error) {
	data := new([]string)
	var reqErr error
	// TODO: need to lookup the region and map it to "americas", "asia", or "EUROPE"
	// TODO: add region/token to the service struct
	resp, err := t.sling.Get("https://americas.api.riotgames.com/tft/match/v1/matches/by-puuid/"+encryptedPUUID+"/ids").Receive(data, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *data, resp, reqErr
}
