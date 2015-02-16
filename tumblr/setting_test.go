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

func CreateTestReblogData() (int64, string){
	/*
	set reblog post id
	*/

	post_id := 0
	reblog_key := ""

	if post_id == 0 {
		post_id = int64(os.Getenv("TUMBLR_REBLOG_POST_ID"))
	}

	if reblog_key == "" {
		cret = os.Getenv("TUMBLR_ACCESS_TOKEN_SECRET")
	}

	return post_id, reblog_key
}
