package tumblr

import (
	"github.com/mrjones/oauth"
	"net/http"
)

const (
	RequestTokenUrl   = "http://www.tumblr.com/oauth/request_token"
	AuthorizeTokenUrl = "http://www.tumblr.com/oauth/authorize"
	AccessTokenUrl    = "http://www.tumblr.com/oauth/access_token"
)

type Client struct {
	consumerKey string
	oauthClient *oauth.Consumer
	token       *oauth.AccessToken
}

func NewClient(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	s := oauth.ServiceProvider{
		RequestTokenUrl:   RequestTokenUrl,
		AuthorizeTokenUrl: AuthorizeTokenUrl,
		AccessTokenUrl:    AccessTokenUrl,
	}

	consumer := oauth.NewConsumer(consumerKey, consumerSecret, s)

	token := &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	return &Client{
		consumerKey: consumerKey,
		oauthClient: consumer,
		token:       token,
	}
}

func (c *Client) Get(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.oauthClient.Get(url, userParams, c.token)
}

func (c *Client) Post(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.oauthClient.Post(url, userParams, c.token)
}

func (c *Client) GetConsumerKey() string {
	return c.consumerKey
}
