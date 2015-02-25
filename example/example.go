package main

import (
	"../tumblr"
	"fmt"
)

func main() {
	consumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessTokenSecret := ""

	t := tumblr.New(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	blogApi := t.NewBlogApi("ota42y.tumblr.com")

	params := make(map[string]string)
	params["offset"] = "1"
	params["limit"] = "1"

	meta, posts, err := blogApi.Photo(&params)
	if err == nil {
		fmt.Println(meta)
		post := (*posts)[0]

		fmt.Println(post.Caption)
		fmt.Println(post.Photos[0].AltSizes[0].Url)
	}
}
