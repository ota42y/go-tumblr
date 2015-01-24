package tumblr

import (
	"os"
)

func CreateTestTumblr() *Tumblr {
	consumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessTokenSecret := ""

	if consumerKey == "" {
		consumerKey = os.Getenv("TUMBLR_CONSUMER_KEY")
	}

	if consumerSecret == "" {
		consumerSecret = os.Getenv("TUMBLR_CONSUMER_SECRET")
	}

	if accessToken == "" {
		accessToken = os.Getenv("TUMBLR_ACCESS_TOKEN")
	}

	if accessTokenSecret == "" {
		accessTokenSecret = os.Getenv("TUMBLR_ACCESS_TOKEN_SECRET")
	}

	return New(consumerKey, consumerSecret, accessToken, accessTokenSecret)
}
