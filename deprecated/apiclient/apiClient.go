// apiClient.go
// Jamf Pro and Jamf Pro Classic API Client

package apiClient

import (
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	uriAuthToken      = "/api/v1/auth/token"
	uriOauthAuthToken = "/api/oauth/token"
	refreshBuffer     = 10 * time.Second // Default refresh buffer to account for clock skew/delays. Can be overridden
)

type Config struct {
	AuthMethod       Authenticator
	URL              string
	HTTPClient       *http.Client
	HttpRetryTimeout time.Duration
	ExtraHeader      map[string]string
	RefreshBuffer    time.Duration
}

type Authenticator interface {
	Authenticate(c *Client) error
}

type BasicAuthConfig struct {
	Username string
	Password string
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

func (o OAuthConfig) Authenticate(c *Client) error {
	c.clientID = o.ClientID
	c.clientSecret = o.ClientSecret
	c.isUsingOAuth = true // Set this flag when using OAuth
	return c.refreshOAuthToken()
}

func (b BasicAuthConfig) Authenticate(c *Client) error {
	c.username = b.Username
	c.password = b.Password
	c.isUsingOAuth = false // Reset this flag when using basic auth
	return c.refreshAuthToken()
}

type Client struct {
	username, password     string
	clientID, clientSecret string
	isUsingOAuth           bool
	url                    string
	refreshBuffer          time.Duration
	token                  *string
	tokenExpiration        *time.Time
	tokenMutex             sync.Mutex

	HttpClient       *http.Client
	HttpRetryTimeout time.Duration
	ExtraHeader      map[string]string
}

type responseAuthToken struct {
	Token   *string    `json:"token,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
}

type responseOAuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (c *Client) setTokenAndExpiration(token *string, expiration *time.Time) {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	c.token = token
	c.tokenExpiration = expiration
}

func (c *Client) getTokenExpiration() *time.Time {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	return c.tokenExpiration
}

func (c *Client) refreshAuthToken() error {
	expiration := c.getTokenExpiration()
	if expiration != nil {
		if expiration.Add(-c.refreshBuffer).After(time.Now()) {
			return nil
		}
	}

	var out responseAuthToken
	if err := c.DoRequest("POST", uriAuthToken, nil, nil, &out); err != nil {
		return errors.Wrap(err, "Failed to refresh auth token")
	}
	c.setTokenAndExpiration(out.Token, out.Expires)

	return nil
}

func (c *Client) refreshOAuthToken() error {
	expiration := c.getTokenExpiration()
	if expiration != nil {
		if expiration.Add(-c.refreshBuffer).After(time.Now()) {
			return nil
		}
	}

	data := map[string]string{
		"client_id":     c.clientID,
		"grant_type":    "client_credentials",
		"client_secret": c.clientSecret,
	}

	var tokenResponse responseOAuthToken
	if err := c.DoRequest("POST", uriOauthAuthToken, data, nil, &tokenResponse); err != nil {
		return errors.Wrap(err, "Failed to refresh OAuth token")
	}

	expiry := time.Now().Add(time.Second * time.Duration(tokenResponse.ExpiresIn-1))
	c.setTokenAndExpiration(&tokenResponse.AccessToken, &expiry)

	return nil
}

func NewClient(cfg Config) (*Client, error) {
	// Default to 10 seconds if not provided
	if cfg.RefreshBuffer == 0 {
		cfg.RefreshBuffer = 10 * time.Second
	}

	c := &Client{
		url:              cfg.URL,
		HttpClient:       cfg.HTTPClient,
		HttpRetryTimeout: cfg.HttpRetryTimeout,
		ExtraHeader:      cfg.ExtraHeader,
		refreshBuffer:    cfg.RefreshBuffer,
	}

	if err := cfg.AuthMethod.Authenticate(c); err != nil {
		return nil, errors.Wrap(err, "Failed to authenticate client")
	}

	return c, nil
}
