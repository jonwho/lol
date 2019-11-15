package lol

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

const (
	baseURL            = "api.riotgames.com/"
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

// ClientOption is a func that operates on *Client
type ClientOption func(*Client) error

// Client API struct to League of Legends
type Client struct {
	Token, Region string
	sling         *sling.Sling
	httpClient    *http.Client
	*LOL
	*TFT
}

// NewClient returns interface to League of Legends API
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
	cli.LOL = NewLOL(cli.sling)
	cli.TFT = NewTFT(cli.sling)

	return cli, nil
}

// WithToken set the client token
func WithToken(token string) ClientOption {
	return func(c *Client) error {
		c.Token = token
		return nil
	}
}

// WithRegion set the client region
func WithRegion(region string) ClientOption {
	return func(c *Client) error {
		c.Region = region
		return nil
	}
}

// WithHTTPClient set the client http.Client and the sling http.Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		c.sling.Client(httpClient)
		return nil
	}
}
