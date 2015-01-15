package main

import (
	"fmt"
	"log"
	"github.com/mrjones/oauth"
)

func main() {
	consumerKey := ""
	consumerSecret := ""

	client := tumblr.New(consumerKey, consumerSecret)

	fmt.Println("(1) Go to: " + client.Client.Test())
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accessToken)
}
