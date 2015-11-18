package tumblr

import (
	"github.com/mrjones/oauth"
	"os"
	"strconv"
)

func CreateTestTumblr() *Tumblr {
	consumerKey := ""
	consumerSecret := ""

	if consumerKey == "" {
		consumerKey = os.Getenv("TUMBLR_CONSUMER_KEY")
	}

	if consumerSecret == "" {
		consumerSecret = os.Getenv("TUMBLR_CONSUMER_SECRET")
	}

	return New(consumerKey, consumerSecret)
}

func CreateTestAPI(c *Client) *BlogAPI {
	accessToken := ""
	accessTokenSecret := ""

	if accessToken == "" {
		accessToken = os.Getenv("TUMBLR_ACCESS_TOKEN")
	}

	if accessTokenSecret == "" {
		accessTokenSecret = os.Getenv("TUMBLR_ACCESS_TOKEN_SECRET")
	}

	token := &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	return NewblogAPI("ota42y.tumblr.com", c, token)
}

func CreateTestReblogData() (int64, string) {
	/*
		set reblog post id
	*/

	postID := int64(0)
	reblogKey := ""

	if postID == 0 {
		postID, _ = strconv.ParseInt(os.Getenv("TUMBLR_REBLOG_POST_ID"), 10, 64)
	}

	if reblogKey == "" {
		reblogKey = os.Getenv("TUMBLR_REBLOG_KEY")
	}

	return postID, reblogKey
}
