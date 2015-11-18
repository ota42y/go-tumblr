package tumblr

import (
	"github.com/mrjones/oauth"
)

// Tumblr is tumblr client struct
type Tumblr struct {
	Client *Client
}

// New return Tumblr
func New(consumerKey string, consumerSecret string) *Tumblr {
	client := NewClient(consumerKey, consumerSecret)

	return &Tumblr{Client: client}
}

// NewblogAPI return blogAPI
func (t *Tumblr) NewblogAPI(host string, token *oauth.AccessToken) *BlogAPI {
	return NewblogAPI(host, t.Client, token)
}
