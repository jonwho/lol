package lol

import (
	"net/http"
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

// Client API struct to League of Legends
type Client struct {
	Token, Region string
	sling         *sling.Sling
	httpClient    *http.Client
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

func (c *Client) AllChampionMastery(encryptedSummonerID string) interface{} {
	return nil
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

func (c *Client) SummonerByName(name string) (*SummonerDTO, *http.Response, error) {
	sd := new(SummonerDTO)
	var reqErr error
	resp, err := c.sling.Get("lol/summoner/v4/summoners/by-name/"+name).Receive(sd, reqErr)
	if err != nil {
		return nil, resp, err
	}
	return sd, resp, reqErr
}
