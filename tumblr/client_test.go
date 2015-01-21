package tumblr

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := CreateTestTumblr().Client

	if client == nil {
		t.Errorf("got %v\nwant %v", client, nil)
	}
}
