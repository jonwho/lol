package lol

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

const (
	baseURL            = "api.riotgames.com/lol"
	defaultRegion      = "na"
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

// Client API struct to League of Legends
type Client struct {
	Token, Region string
	slinger       *sling.Sling
	httpClient    *http.Client
}

// ClientOption is a func that operates on *Client
type ClientOption func(*Client) error

// New returns interface to League of Legends API
func NewClient(token string, options ...ClientOption) (*Client, error) {
	cli := &Client{
		Token:   token,
		Region:  defaultRegion,
		slinger: sling.New().Base(baseURL),
	}

	for _, option := range options {
		if err := option(cli); err != nil {
			return nil, err
		}
	}

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
		c.slinger.Client(httpClient)
		return nil
	}
}
