package lol

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
)

var (
	testToken  = os.Getenv("RIOT_API_KEY")
	httpClient *http.Client
)

func TestNewClient(t *testing.T) {
	cli, err := NewClient("test_key")
	if err != nil {
		t.Error(err)
	}
	expected := "test_key"
	actual := cli.Token
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
	expected = "na1"
	actual = cli.Region
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	cli, err = NewClient("to_be_overwritten", WithToken("foobar"))
	if err != nil {
		t.Error(err)
	}
	expected = "foobar"
	actual = cli.Token
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	cli, err = NewClient("to_be_overwritten", WithRegion("mynewregion"))
	if err != nil {
		t.Error(err)
	}
	expected = "mynewregion"
	actual = cli.Region
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestChampionRotations(t *testing.T) {
	rec, err := recorder.New("cassettes/champion-v3/champion-rotations")
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
	}

	championInfo, resp, err := cli.ChampionRotations()
	log.Println(resp.Status)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := false
	actual := len(championInfo.FreeChampionIDs) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSummonerByName(t *testing.T) {
	rec, err := recorder.New("cassettes/summoner-v4/summoner-by-name")
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
	}

	sd, resp, err := cli.SummonerByName("ilikeduck")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := 70
	actual := sd.SummonerLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
	}
}

func matchWithoutToken(req *http.Request, i cassette.Request) bool {
	u := req.URL
	q := u.Query()
	q.Del("api_key")
	u.RawQuery = q.Encode()
	req.URL = u
	return u.String() == i.URL
}

func removeToken(i *cassette.Interaction) error {
	// remove from URL
	u, err := url.Parse(i.Request.URL)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Del("api_key")
	u.RawQuery = q.Encode()
	i.Request.URL = u.String()

	// remove from JSON request body
	originalBody := []byte(i.Request.Body)
	var unmarshalBody map[string]interface{}
	if err = json.Unmarshal(originalBody, &unmarshalBody); err != nil {
		// try to unmarshal response body to JSON
		// NOP if error
	}
	delete(unmarshalBody, "api_key")
	bodyWithoutToken, err := json.Marshal(unmarshalBody)
	i.Request.Body = string(bodyWithoutToken)

	// remove from header
	delete(i.Request.Headers, "X-Riot-Token")
	return nil
}
