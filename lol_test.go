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
	testToken           = os.Getenv("RIOT_API_KEY")
	httpClient          *http.Client
	encryptedSummonerID = "1NgBFb-1WXj-ku_Fym3BQF1FxXUz9xrvpuIPVnSdvo6KjHo"
	grandmasterLeagueID = "00d07caf-539b-346a-a4f8-fdb57ab31aa4"
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

func TestAllChampionMastery(t *testing.T) {
	rec, err := recorder.New("cassettes/champion-mastery-v4/all-champion-mastery")
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

	dtosPointer, resp, err := cli.AllChampionMastery(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	dtos := *dtosPointer
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestChampionMastery(t *testing.T) {
	rec, err := recorder.New("cassettes/champion-mastery-v4/champion-mastery")
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

	dto, resp, err := cli.ChampionMastery(encryptedSummonerID, "39")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := 7
	actual := dto.ChampionLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
	}
}

func TestMasteryScore(t *testing.T) {
	rec, err := recorder.New("cassettes/champion-mastery-v4/mastery-score")
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

	score, resp, err := cli.MasteryScore(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := 99
	actual := score
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
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

func TestLeagueExpEntries(t *testing.T) {
	rec, err := recorder.New("cassettes/league-exp-v4/league-exp-entries")
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

	dtos, resp, err := cli.LeagueExpEntries("RANKED_SOLO_5x5", "CHALLENGER", "I", &LeagueExpEntriesParams{Page: "1"})
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestChallengerLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/challenger-leagues")
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

	dto, resp, err := cli.ChallengerLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := "Syndra's Masterminds"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestEntriesBySummoner(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/entries-by-summoner")
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

	dtos, resp, err := cli.EntriesBySummoner(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestEntries(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/entries")
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

	dtos, resp, err := cli.Entries("RANKED_SOLO_5x5", "DIAMOND", "I", &EntriesParams{Page: "1"})
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestGrandmasterLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/grandmaster-leagues")
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

	dto, resp, err := cli.GrandmasterLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := "Hecarim's Duelists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/leagues")
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

	dto, resp, err := cli.Leagues(grandmasterLeagueID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := "Hecarim's Duelists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestMasterLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/league-v4/master-leagues")
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

	dto, resp, err := cli.MasterLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
	}
	if err != nil {
		t.Error(err)
	}
	expected := "Jarvan IV's Elementalists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
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

func TestSummonerByPUUID(t *testing.T) {
	rec, err := recorder.New("cassettes/summoner-v4/summoner-by-puuid")
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

	sd, resp, err := cli.SummonerByPUUID("HldoCYMHNm27w37qJCfk5d20dB5uGma7oNuBVoZ01n3do7fMLW7ubao6SDeVAqTd9ieB5orqXvwHsQ")
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
