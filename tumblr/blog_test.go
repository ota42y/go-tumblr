package tumblr

import (
	"testing"
)

func TestNew(t *testing.T) {
	client := NewClient("", "", "", "")
	actual := NewBlog("scipsy.tumblr.com", client)
	if actual == nil {
		t.Errorf("got %v\nwant %v", actual, nil)
	}
}

func TestInfo(t *testing.T) {
	client := NewClient("", "", "", "")
	blog := NewBlog("scipsy.tumblr.com", client)
	actual, err := blog.Info()

	// debug
	t.Errorf("got %v", actual)

	if actual == nil{
		t.Errorf("got %v\nwant %v error is %v", actual, nil, err)
	}

	validCode := 200
	if actual.StatusCode != validCode{
		t.Errorf("%v is not response code %v", actual.StatusCode, validCode)
	}
}
