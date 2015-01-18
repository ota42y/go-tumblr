package tumblr

type Client struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func NewClient(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	return &Client{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken: accessToken,
		AccessTokenSecret: accessTokenSecret,
	}
}
