package lol

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
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
