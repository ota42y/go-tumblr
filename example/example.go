package main

import (
	"fmt"

	"github.com/ota42y/go-tumblr/tumblr"
	"github.com/typester/go-pit"
)

var tumblrDomain string = "ota42y.tumblr.com"

func main() {
	p, err := pit.Get(tumblrDomain, pit.Requires{
		"consumer_key":        "",
		"consumer_secret":    "",
		"access_token":    "",
		"access_token_secret": "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}


	consumerKey := (*p)["consumer_key"]
	consumerSecret := (*p)["consumer_secret"]
	accessToken := (*p)["access_token"]
	accessTokenSecret := (*p)["access_token_secret"]

	t := tumblr.New(consumerKey, consumerSecret)

	, accessToken, accessTokenSecret)
	blogAPI := t.NewblogAPI(tumblrDomain)

	params := make(map[string]string)
	params["offset"] = "1"
	params["limit"] = "1"

	meta, posts, err := blogAPI.Photo(&params)
	if err == nil {
		fmt.Println(meta)
		post := (*posts)[0]

		fmt.Println(post.Caption)
		fmt.Println(post.Photos[0].AltSizes[0].URL)
	}
}
