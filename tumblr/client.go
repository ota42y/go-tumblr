package tumblr

type Client struct {
	consumerKey string
	consumerSecret string
	accessToken string
	accessTokenSecret string
}

func NewClient(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Client {
	return &Client{
		consumerKey: consumerKey,
		consumerSecret: consumerSecret,
		accessToken: accessToken,
		accessTokenSecret: accessTokenSecret,
	}
}
