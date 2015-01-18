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
	actual := blog.Info()

	if actual == nil{
		t.Errorf("got %v\nwant %v", actual, nil)
	}


}
