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
	summonerName        = "ilikeduck"
	encryptedAccountID  = "L019WecOvXAAA7U2pplSIFOUjOvleGyX_9X_p2Al7J007A"
	encryptedSummonerID = "1NgBFb-1WXj-ku_Fym3BQF1FxXUz9xrvpuIPVnSdvo6KjHo"
	encryptedPUUID      = "HldoCYMHNm27w37qJCfk5d20dB5uGma7oNuBVoZ01n3do7fMLW7ubao6SDeVAqTd9ieB5orqXvwHsQ"
	grandmasterLeagueID = "00d07caf-539b-346a-a4f8-fdb57ab31aa4"
	matchID             = "3198831326"
)

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

func TestAllChampionMastery(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/champion-mastery-v4/all-champion-mastery")
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

	dtosPointer, resp, err := cli.AllChampionMastery(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	dtos := *dtosPointer
	expected := false
	actual := len(dtos) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestChampionMastery(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/champion-mastery-v4/champion-mastery")
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

	dto, resp, err := cli.ChampionMastery(encryptedSummonerID, "39")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 7
	actual := dto.ChampionLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestMasteryScore(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/champion-mastery-v4/mastery-score")
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

	score, resp, err := cli.MasteryScore(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 99
	actual := score
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestChampionRotations(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/champion-v3/champion-rotations")
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

	championInfo, resp, err := cli.ChampionRotations()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(championInfo.FreeChampionIDs) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestLeagueExpEntries(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-exp-v4/league-exp-entries")
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

	dtos, resp, err := cli.LeagueExpEntries("RANKED_SOLO_5x5", "CHALLENGER", "I", &LeagueExpEntriesParams{Page: "1"})
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

func TestChallengerLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/challenger-leagues")
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

	dto, resp, err := cli.ChallengerLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Syndra's Masterminds"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestEntriesBySummoner(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/entries-by-summoner")
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

	dtos, resp, err := cli.LOL.EntriesBySummoner(encryptedSummonerID)
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

func TestEntries(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/entries")
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

	dtos, resp, err := cli.LOL.Entries("RANKED_SOLO_5x5", "DIAMOND", "I", &EntriesParams{Page: "1"})
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

func TestGrandmasterLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/grandmaster-leagues")
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

	dto, resp, err := cli.GrandmasterLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Hecarim's Duelists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/leagues")
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

	dto, resp, err := cli.LOL.Leagues(grandmasterLeagueID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Hecarim's Duelists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestMasterLeagues(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/league-v4/master-leagues")
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

	dto, resp, err := cli.MasterLeagues("RANKED_SOLO_5x5")
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Jarvan IV's Elementalists"
	actual := dto.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestStatus(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/lol-status-v3/status")
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

	status, resp, err := cli.Status()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := "North America"
	actual := status.Name
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
		return
	}
}

func TestMatches(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/match-v4/matches")
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

	dto, resp, err := cli.Matches(matchID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 13
	actual := dto.SeasonID
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestMatchlists(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/match-v4/matchlists")
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

	ml, resp, err := cli.Matchlists(encryptedAccountID, nil)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(ml.Matches) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestTimelines(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/match-v4/timelines")
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

	tl, resp, err := cli.Timelines(matchID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(tl.Frames) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestActiveGames(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/spectator-v4/active-games")
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

	ag, resp, err := cli.ActiveGames(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	var expected int64
	expected = 3206061846
	actual := ag.GameID
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestFeaturedGames(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/spectator-v4/featured-games")
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

	fg, resp, err := cli.FeaturedGames()
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := false
	actual := len(fg.GameList) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
		return
	}
}

func TestSummonerByAccount(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/summoner-v4/summoner-by-account")
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

	sd, resp, err := cli.SummonerByAccount(encryptedAccountID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 73
	actual := sd.SummonerLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}
func TestSummonerByName(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/summoner-v4/summoner-by-name")
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

	sd, resp, err := cli.SummonerByName(summonerName)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 70
	actual := sd.SummonerLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestSummonerByPUUID(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/summoner-v4/summoner-by-puuid")
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

	sd, resp, err := cli.SummonerByPUUID(encryptedPUUID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 70
	actual := sd.SummonerLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}

func TestSummonerByID(t *testing.T) {
	rec, err := recorder.New("cassettes/lol/summoner-v4/summoner-by-id")
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

	sd, resp, err := cli.SummonerByID(encryptedSummonerID)
	if resp.StatusCode != 200 {
		t.Errorf("\nExpected: 200 status code\nActual: %d status code", resp.StatusCode)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	expected := 73
	actual := sd.SummonerLevel
	if expected != actual {
		t.Errorf("\nExpected: %d\nActual: %d\n", expected, actual)
		return
	}
}
