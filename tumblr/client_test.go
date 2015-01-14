package tumblr

import (
	"testing"
)

func TestNew(t *testing.T) {
	consumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessTokenSecret := ""

	client := NewClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	if client == nil {
		t.Errorf("got %v\nwant %v", client, nil)
	}
}
