package tumblr

import (
	"github.com/mrjones/oauth"
	"net/http"
)

const (
	requestTokenURL   = "http://www.tumblr.com/oauth/request_token"
	authorizeTokenURL = "http://www.tumblr.com/oauth/authorize"
	accessTokenURL    = "http://www.tumblr.com/oauth/access_token"
)

// Client is tumblr api client
type Client struct {
	consumerKey string
	oauthClient *oauth.Consumer
}

// NewClient return Client
func NewClient(consumerKey string, consumerSecret string) *Client {
	s := oauth.ServiceProvider{
		RequestTokenUrl:   requestTokenURL,
		AuthorizeTokenUrl: authorizeTokenURL,
		AccessTokenUrl:    accessTokenURL,
	}

	consumer := oauth.NewConsumer(consumerKey, consumerSecret, s)

	return &Client{
		consumerKey: consumerKey,
		oauthClient: consumer,
	}
}

// Get send data by GET method
func (c *Client) Get(url string, userParams map[string]string, token *oauth.AccessToken) (resp *http.Response, err error) {
	return c.oauthClient.Get(url, userParams, token)
}

// Post send data by POST method
func (c *Client) Post(url string, userParams map[string]string, token *oauth.AccessToken) (resp *http.Response, err error) {
	return c.oauthClient.Post(url, userParams, token)
}

// GetConsumerKey return consumer key
func (c *Client) GetConsumerKey() string {
	return c.consumerKey
}

// GetRequestTokenAndURL get request token and authorize url
func (c *Client) GetRequestTokenAndURL(callbackURL string) (rtoken *oauth.RequestToken, loginURL string, err error) {
	return c.oauthClient.GetRequestTokenAndUrl(callbackURL)
}

// AuthorizeToken get tumblr auth token
func (c *Client) AuthorizeToken(rtoken *oauth.RequestToken, verificationCode string) (atoken *oauth.AccessToken, err error) {
	return c.oauthClient.AuthorizeToken(rtoken, verificationCode)
}
