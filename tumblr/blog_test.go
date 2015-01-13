package tumblr

import (
	"testing"
)

func TestNew(t *testing.T) {
	actual := NewBlog("scipsy.tumblr.com")
	if actual == nil {
		t.Errorf("got %v\nwant %v", actual, nil)
	}
}

func TestInfo(t *testing.T) {
	blog := NewBlog("scipsy.tumblr.com")
	actual := blog.Info()

	if actual == nil{
		t.Errorf("got %v\nwant %v", actual, nil)
	}


}
