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

func CreateTestApi(c *Client) *BlogApi {
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

	return NewBlogApi("ota42y.tumblr.com", c, token)
}

func CreateTestReblogData() (int64, string) {
	/*
		set reblog post id
	*/

	post_id := int64(0)
	reblog_key := ""

	if post_id == 0 {
		post_id, _ = strconv.ParseInt(os.Getenv("TUMBLR_REBLOG_POST_ID"), 10, 64)
	}

	if reblog_key == "" {
		reblog_key = os.Getenv("TUMBLR_REBLOG_KEY")
	}

	return post_id, reblog_key
}
