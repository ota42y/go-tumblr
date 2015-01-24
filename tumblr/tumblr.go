package tumblr

type Tumblr struct {
	Client *Client
}

func New(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Tumblr {
	client := NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)

	return &Tumblr{Client: client}
}
