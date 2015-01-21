package tumblr

import (
	"testing"
)

func TestNewBlogApi(t *testing.T) {
	client := CreateTestTumblr().Client
	actual := NewBlogApi("scipsy.tumblr.com", client)
	if actual == nil {
		t.Errorf("got %v\nwant %v", actual, nil)
	}
}

func TestInfo(t *testing.T) {
	client := CreateTestTumblr().Client
	blogApi := NewBlogApi("scipsy.tumblr.com", client)
	meta, blog, err := blogApi.Info()

	if err != nil{
		t.Errorf("response error%v\n", err)
	}

	if meta == nil{
		t.Errorf("got %v\nwant %v error is %v\n", meta, nil, err)
	}

	if blog == nil{
		t.Errorf("got %v\nwant %v error is %v\n", blog, nil, err)
	}

	validCode := 200
	if meta.Status != validCode{
		t.Errorf("%v is not response code %v", meta.Status, validCode)
	}

	if blog.Title == ""{
		t.Errorf("%v is not valid response", blog.Title)
	}
}
