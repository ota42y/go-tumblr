package tumblr

import (
	"github.com/mrjones/oauth"
)

type Tumblr struct {
	Client *Client
}

func New(consumerKey string, consumerSecret string) *Tumblr {
	client := NewClient(consumerKey, consumerSecret)

	return &Tumblr{Client: client}
}

func (t *Tumblr) NewBlogApi(host string, token *oauth.AccessToken) *BlogApi {
	return NewBlogApi(host, t.Client, token)
}
