package tumblr

import (
	"testing"
)

func TestNewBlogApi(t *testing.T) {
	client := CreateTestTumblr().Client
	actual := CreateTestApi(client)
	if actual == nil {
		t.Errorf("got %v\nwant %v", actual, nil)
	}
}

func TestInfo(t *testing.T) {
	client := CreateTestTumblr().Client
	blogApi := CreateTestApi(client)
	meta, blog, err := blogApi.Info()

	if err != nil {
		t.Errorf("response error%v\n", err)
	}

	if meta == nil {
		t.Errorf("got %v\nwant %v error is %v\n", meta, nil, err)
	}

	if blog == nil {
		t.Errorf("got %v\nwant %v error is %v\n", blog, nil, err)
	}

	validCode := 200
	if meta.Status != validCode {
		t.Errorf("%v is not response code %v", meta.Status, validCode)
	}

	if blog.Title == "" {
		t.Errorf("%v is not valid response", blog.Title)
	}
}

func TestPosts(t *testing.T) {
	client := CreateTestTumblr().Client
	blogApi := CreateTestApi(client)
	meta, posts, err := blogApi.Posts("")

	if err != nil {
		t.Errorf("response error%v\n", err)
	}

	if meta == nil {
		t.Errorf("meat is nil")
	}

	if posts == nil {
		t.Errorf("response is nil")
	}

	if len(*posts) == 0 {
		t.Errorf("no posts")
	}
}

func TestPhotos(t *testing.T) {
	client := CreateTestTumblr().Client
	blogApi := CreateTestApi(client)
	meta, posts, err := blogApi.Photo()

	if err != nil {
		t.Errorf("response error%v\n", err)
	}

	if meta == nil {
		t.Errorf("meat is nil")
	}

	if posts == nil {
		t.Errorf("response is nil")
	}

	if len(*posts) == 0 {
		t.Errorf("no posts %v", meta)
	}

	if len((*posts)[0].Photos) == 0 {
		t.Errorf("photos isn't exist")
	}

	if len((*posts)[0].Photos[0].AltSizes) == 0 {
		t.Errorf("alt sizes isn't exist")
	}
}

func TestQuote(t *testing.T) {
	client := CreateTestTumblr().Client
	blogApi := CreateTestApi(client)
	meta, posts, err := blogApi.Quote()

	if err != nil {
		t.Errorf("response error%v\n", err)
	}

	if meta == nil {
		t.Errorf("meat is nil")
	}

	if posts == nil {
		t.Errorf("response is nil")
	}

	if len(*posts) == 0 {
		t.Errorf("no posts %v", meta)
		return
	}

	post := (*posts)[0]
	if post.Text == "" {
		t.Errorf("text isn't exist")
	}

	if post.Source == "" {
		t.Errorf("source isn't exist")
	}
}

func TestReblog(t *testing.T) {
	id, reblog_key := CreateTestReblogData()

	client := CreateTestTumblr().Client
	blogApi := CreateTestApi(client)
	meta, id, err := blogApi.Reblog(id, reblog_key, "test")

	if err != nil {
		t.Errorf("response error%v\n", err)
	}

	if meta == nil {
		t.Errorf("response is nil")
	}

	if id == 0 {
		t.Errorf("id is 0")
	}
}
