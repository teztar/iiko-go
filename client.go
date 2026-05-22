package iiko

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	// Channel quit is used to notify that we should stop our JWT-refresh token Ticker.
	quit chan struct{}

	baseURL  string
	apiLogin string

	// appId and clientSecret belong to the new iiko authorization scheme
	// (mandatory from 2026-06-01). Empty => fall back to old apiLogin-only.
	appId        string
	clientSecret string

	// tokenMu guards token against the data race between refreshTokenByInterval
	// (writer) and request builders (readers).
	tokenMu sync.RWMutex
	token   string

	httpClient           *http.Client
	timeout              time.Duration
	refreshTokenInterval time.Duration
}

// ClientOption customizes a Client at construction time.
type ClientOption func(*Client)

// WithApp enables the new iiko authorization scheme by attaching the
// integration's appId and clientSecret to every access_token request.
func WithApp(appId, clientSecret string) ClientOption {
	return func(c *Client) {
		c.appId = appId
		c.clientSecret = clientSecret
	}
}

// SetTimeout sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetTimeout(t time.Duration) {
	c.timeout = t
}

// SetRefreshTokenInterval sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetRefreshTokenInterval(t time.Duration) {
	c.refreshTokenInterval = t
}

// SetHTTPClient sets a custom http.Client for making API request to iikoCloud.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// getToken returns the current access token in a thread-safe way.
func (c *Client) getToken() string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.token
}

// setToken stores a new access token in a thread-safe way.
func (c *Client) setToken(token string) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	c.token = token
}

// refreshToken fetches a fresh access token from iikoCloud and stores it.
// On any error it keeps the previous token untouched so in-flight requests
// keep working until the next successful refresh.
func (c *Client) refreshToken() error {
	resp, err := c.accessToken(&AccessTokenRequest{
		ApiLogin:     c.apiLogin,
		AppId:        c.appId,
		ClientSecret: c.clientSecret,
	})
	if err != nil {
		return err
	}
	if resp == nil || resp.Token == "" {
		return errors.New("iiko: empty access token in response")
	}
	c.setToken(resp.Token)
	return nil
}

func (c *Client) refreshTokenByInterval() {
	ticker := time.NewTicker(c.refreshTokenInterval)

	for {
		select {
		case <-ticker.C:
			// Ignore the error on purpose: keep the old (still valid) token and
			// retry on the next tick. Never panic on a nil response — that used
			// to kill this goroutine and freeze the token forever, causing 401s
			// an hour later.
			_ = c.refreshToken()

		case <-c.quit:
			ticker.Stop()
			return
		}
	}
}

func (c *Client) Close() {
	close(c.quit)
}

func NewClient(apiLogin string, opts ...ClientOption) (*Client, error) {
	client := &Client{
		baseURL:              BaseURL,
		httpClient:           http.DefaultClient,
		apiLogin:             apiLogin,
		timeout:              DefaultTimeout,
		refreshTokenInterval: DefaultRefreshTokenInterval,
		quit:                 make(chan struct{}),
	}

	for _, opt := range opts {
		opt(client)
	}

	if err := client.refreshToken(); err != nil {
		return nil, err
	}

	go client.refreshTokenByInterval()

	return client, nil
}
