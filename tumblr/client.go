package tumblr

import (
	"github.com/mrjones/oauth"
	"net/http"
)

const (
	RequestTokenUrl = "http://www.tumblr.com/oauth/request_token"
	AuthorizeTokenUrl   = "http://www.tumblr.com/oauth/authorize"
	AccessTokenUrl = "http://www.tumblr.com/oauth/access_token"
)

type Client struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string

	oauthClient *oauth.Consumer
	token *oauth.AccessToken
}

func NewClient(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	s := oauth.ServiceProvider{
		RequestTokenUrl: RequestTokenUrl,
		AuthorizeTokenUrl: AuthorizeTokenUrl,
		AccessTokenUrl: AccessTokenUrl,
	}

	consumer := oauth.NewConsumer(consumerKey, consumerSecret, s)

	token := &oauth.AccessToken{
		Token: accessToken,
		Secret: accessTokenSecret,
	}

	return &Client{
		ConsumerKey:       consumerKey,
		ConsumerSecret:    consumerSecret,
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
		oauthClient: consumer,
		token: token,
	}
}

func (c *Client) Get(url string, userParams map[string]string) (resp *http.Response, err error) {
	return c.oauthClient.Get(url, userParams, c.token)
}
