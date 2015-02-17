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

	meta, posts, err := blogApi.Photo()
	if err == nil {
		fmt.Println(meta)
		post := (*posts)[0]

		fmt.Println(post.Caption)
		fmt.Println(post.Photos[0].AltSizes[0].Url)
	}
}
