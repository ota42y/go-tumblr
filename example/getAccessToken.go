package main

// This file base on a oauth example.
// https://github.com/mrjones/oauth

import (
	"../tumblr"
	"fmt"
)

func main() {
	consumerKey := ""
	consumerSecret := ""

	t := tumblr.New(consumerKey, consumerSecret)
	client := t.Client

	requestToken, url, err := client.GetRequestTokenAndUrl("http://localhost")
	if err != nil {
		panic(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := client.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		panic(err)
	}

	fmt.Println(accessToken.Token)
	fmt.Println(accessToken.Secret)

	blogApi := t.NewBlogApi("ota42y.tumblr.com", accessToken)

	meta, posts, err := blogApi.Photo()
	if err == nil {
		fmt.Println(meta)
		post := (*posts)[0]

		fmt.Println(post.Caption)
		fmt.Println(post.Photos[0].AltSizes[0].Url)
	}
}
