package lol

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

var (
	tftSummonerName        = "bnage"
	tftEncryptedAccountID  = "7cwQwRgPliNgswYKofJlT-xM3WHd_IPWnK-uak1Q5tjFKto"
	tftEncryptedSummonerID = "mZB3KRfmKzq0uo1LA8yVdClbaDAfPev_GNBaocjYcHpt6Ik"
	tftEncryptedPUUID      = "yb7CinbPRVCa25cTwjeFxpBVpsggU0c2emAl7Rfi0LcCTUppGb0q393un2JsgHpKGJGb7sDelhNZug"
	tftLeagueID            = "302f7830-005c-11ea-9566-da80b681e2c4"
)

func TestChallenger(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/challenger")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	dto, resp, err := cli.Challenger()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Trundle's Stalkers"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestTFTEntriesBySummoner(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/entries-by-summoner")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	dtos, resp, err := cli.TFT.EntriesBySummoner(tftEncryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestTFTEntries(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/entries")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	tftTier := "DIAMOND"
	tftDivision := "I"
	dtos, resp, err := cli.TFT.Entries(tftTier, tftDivision, &EntriesParams{Page: "2"})
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestGrandmaster(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/grandmaster")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	dto, resp, err := cli.Grandmaster()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Sona's Zealots"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestTFTLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/leagues")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	dto, resp, err := cli.TFT.Leagues(tftLeagueID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Ashe's Battlemasters"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestMaster(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/league-v1/master")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	dto, resp, err := cli.Master()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Lux's Dawnbringers"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestMatchesByPUUID(t *testing.T) {
	rec, err := recorder.New("cassettes/tft/match-v1/matches-by-puuid")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, WithHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
		return
	}

	data, resp, err := cli.MatchesByPUUID(tftEncryptedPUUID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(data) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}
