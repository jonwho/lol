package lol

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dghubble/sling"
)

const (
	baseURL            = "api.riotgames.com/lol"
	defaultRegion      = "na1"
	maxIdleConnections = 10
	requestTimeout     = 5
)

var (
	// DefaultHTTPClient default http client to use
	DefaultHTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
)

type ChampionInfo struct {
	FreeChampionIDs              []int `json:"freeChampionIds"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

type ChampionMasteryDTO struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}

// Client API struct to League of Legends
type Client struct {
	Token, Region string
	sling         *sling.Sling
	httpClient    *http.Client
}

type EntriesParams struct {
	Page string `url:"page,omitempty"`
}

type Incident struct {
	Active    bool      `json:"active"`
	CreatedAt string    `json:"created_at"`
	ID        int       `json:"id"`
	Updates   []Message `json:"updates"`
}

type LeagueEntryDTO struct {
	QueueType    string        `json:"queueType"`
	SummonerName string        `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Wins         int           `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int           `json:"losses"`
	Rank         string        `json:"rank"`
	Tier         string        `json:"tier"`
	Inactive     bool          `json:"inactive"`
	FreshBlood   bool          `json:"freshBlood"`
	LeagueID     string        `json:"leagueId"`
	SummonerID   string        `json:"summonerId"`
	LeaguePoints int           `json:"leaguePoints"`
}

type LeagueExpEntriesParams struct {
	Page string `url:"page,omitempty"`
}

type LeagueItemDTO struct {
	SummonerName string        `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Wins         int           `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int           `json:"losses"`
	FreshBlood   bool          `json:"freshBlood"`
	Inactive     bool          `json:"inactive"`
	Rank         string        `json:"rank"`
	SummonerID   string        `json:"summonerId"`
	LeaguePoints int           `json:"leaguePoints"`
}

type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

type MatchlistDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

type MatchlistsParams struct {
	Champion   []int `url:"champion"`
	Queue      []int `url:"queue"`
	Season     []int `url:"season"`
	EndTime    int   `url:"endTime"`
	BeginTime  int   `url:"beginTime"`
	EndIndex   int   `url:"endIndex"`
	BeginIndex int   `url:"beginIndex"`
}

type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int    `json:"timestamp"`
}

type Message struct {
	Severity     string        `json:"severity"`
	Author       string        `json:"author"`
	CreatedAt    string        `json:"created_at"`
	Translations []Translation `json:"translations"`
	UpdatedAt    string        `json:"updated_at"`
	Content      string        `json:"content"`
	ID           string        `json:"id"`
}

type MiniSeriesDTO struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

type ShardStatus struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Hostname  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

type Translation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// ClientOption is a func that operates on *Client
type ClientOption func(*Client) error

// New returns interface to League of Legends API
func NewClient(token string, options ...ClientOption) (*Client, error) {
	cli := &Client{}
	WithToken(token)(cli)
	WithRegion(defaultRegion)(cli)
	cli.sling = sling.New().Base("https://" + cli.Region + "." + baseURL)
	cli.sling.Set("User-Agent", "jonwho/lol")

	for _, option := range options {
		if err := option(cli); err != nil {
			return nil, err
		}
	}

	cli.sling.Set("X-Riot-Token", cli.Token)

	return cli, nil
}

func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.Token = token
		return nil
	}
}

func WithRegion(region string) ClientOption {
	return func(c *Client) error {
		c.Region = region
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		c.sling.Client(httpClient)
		return nil
	}
}

// AllChampionMastery GET /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerID}
func (c *Client) AllChampionMastery(encryptedSummonerID string) (*[]ChampionMasteryDTO, *http.Response, error) {
	dtos := new([]ChampionMasteryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID).Receive(dtos, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return dtos, resp, reqErr
}

// ChampionMastery GET /lol/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerID}/by-champion/{championID}
func (c *Client) ChampionMastery(encryptedSummonerID, championID string) (*ChampionMasteryDTO, *http.Response, error) {
	dto := new(ChampionMasteryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/champion-mastery/v4/champion-masteries/by-summoner/"+encryptedSummonerID+"/by-champion/"+championID).Receive(dto, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return dto, resp, reqErr
}

// MasteryScore GET /lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerID}
func (c *Client) MasteryScore(encryptedSummonerID string) (int, *http.Response, error) {
	req, err := c.sling.Get("lol/champion-mastery/v4/scores/by-summoner/" + encryptedSummonerID).Request()
	if err != nil {
		return 0, nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, resp, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, resp, err
	}
	score, err := strconv.Atoi(string(bodyBytes))
	if err != nil {
		return 0, resp, err
	}
	return score, resp, err
}

// ChampionRotations GET /lol/platform/v3/champion-rotations
func (c *Client) ChampionRotations() (*ChampionInfo, *http.Response, error) {
	ci := new(ChampionInfo)
	var reqErr error
	resp, err := c.sling.Get("lol/platform/v3/champion-rotations").Receive(ci, reqErr)

	if err != nil {
		return nil, resp, err
	}

	return ci, resp, reqErr
}

// LeagueExpEntries GET /lol/league-exp/v4/entries/{queue}/{tier}/{division}
func (c *Client) LeagueExpEntries(queue, tier, division string, params *LeagueExpEntriesParams) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	endpoint := fmt.Sprintf("lol/league-exp/v4/entries/%s/%s/%s", queue, tier, division)
	resp, err := c.sling.Get(endpoint).QueryStruct(params).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// ChallengerLeagues GET /lol/league/v4/challengerleagues/by-queue/{queue}
func (c *Client) ChallengerLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/challengerleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// EntriesBySummoner GET /lol/league/v4/entries/by-summoner/{encryptedSummonerId}
func (c *Client) EntriesBySummoner(encryptedSummonerID string) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/entries/by-summoner/"+encryptedSummonerID).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// Entries GET /lol/league/v4/entries/{queue}/{tier}/{division}
func (c *Client) Entries(queue, tier, division string, params *EntriesParams) ([]LeagueEntryDTO, *http.Response, error) {
	dtos := new([]LeagueEntryDTO)
	var reqErr error
	endpoint := fmt.Sprintf("lol/league/v4/entries/%s/%s/%s", queue, tier, division)
	resp, err := c.sling.Get(endpoint).Receive(dtos, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return *dtos, resp, reqErr
}

// GrandmasterLeagues GET /lol/league/v4/grandmasterleagues/by-queue/{queue}
func (c *Client) GrandmasterLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/grandmasterleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Leagues GET /lol/league/v4/leagues/{leagueId}
func (c *Client) Leagues(leagueID string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/leagues/"+leagueID).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// MasterLeagues GET /lol/league/v4/masterleagues/by-queue/{queue}
func (c *Client) MasterLeagues(queue string) (*LeagueListDTO, *http.Response, error) {
	dto := new(LeagueListDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/league/v4/masterleagues/by-queue/"+queue).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// Status GET /lol/status/v3/shard-data
func (c *Client) Status() (*ShardStatus, *http.Response, error) {
	shardStatus := new(ShardStatus)
	var reqErr error
	resp, err := c.sling.Get("lol/status/v3/shard-data").Receive(shardStatus, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return shardStatus, resp, reqErr
}

// Matchlists GET /lol/match/v4/matchlists/by-account/{encryptedAccountID}
func (c *Client) Matchlists(encryptedAccountID string, params *MatchlistsParams) (*MatchlistDTO, *http.Response, error) {
	dto := new(MatchlistDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/match/v4/matchlists/by-account/"+encryptedAccountID).QueryStruct(params).Receive(dto, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return dto, resp, reqErr
}

// SummonerByName GET /lol/summoner/v4/summoners/by-name/{summonerName}
func (c *Client) SummonerByName(summonerName string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-name/"+summonerName).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}

// SummonerByPUUID GET /lol/summoner/v4/summoners/by-puuid/{encryptedPUUID}
func (c *Client) SummonerByPUUID(encryptedPUUID string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-puuid/"+encryptedPUUID).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}
