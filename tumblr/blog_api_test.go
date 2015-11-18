package tumblr

import (
	"testing"
)

func TestNewblogAPI(t *testing.T) {
	client := CreateTestTumblr().Client
	actual := CreateTestAPI(client)
	if actual == nil {
		t.Errorf("got %v\nwant %v", actual, nil)
	}
}

func TestInfo(t *testing.T) {
	client := CreateTestTumblr().Client
	blogAPI := CreateTestAPI(client)
	meta, blog, err := blogAPI.Info()

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
	blogAPI := CreateTestAPI(client)
	meta, posts, err := blogAPI.Posts("", nil)

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
	blogAPI := CreateTestAPI(client)
	meta, posts, err := blogAPI.Photo(nil)

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

func TestPhotosWithLimit(t *testing.T) {
	client := CreateTestTumblr().Client
	blogAPI := CreateTestAPI(client)

	meta, posts, err := blogAPI.Posts("", nil)

	if err != nil {
		t.Errorf("response error%v\n", err)
		t.FailNow()
	}

	if meta == nil {
		t.Errorf("meat is nil")
		t.FailNow()
	}

	if posts == nil {
		t.Errorf("response is nil")
		t.FailNow()
	}

	if len(*posts) < 2 {
		t.Errorf("paramater's test need 2 or more post %v", meta)
		t.FailNow()
	}

	params := make(map[string]string)
	params["offset"] = "1"
	params["limit"] = "1"
	meta, offsetPosts, err := blogAPI.Posts("", &params)

	if err != nil {
		t.Errorf("response error%v\n", err)
		t.FailNow()
	}

	if meta == nil {
		t.Errorf("meat is nil")
		t.FailNow()
	}

	if offsetPosts == nil {
		t.Errorf("response is nil")
		t.FailNow()
	}

	if len(*offsetPosts) != 1 {
		t.Errorf("no posts %v", meta)
		t.FailNow()
	}

	if (*posts)[1].ID != (*offsetPosts)[0].ID {
		t.Errorf("offset isn't work")
		t.FailNow()
	}
}
func TestQuote(t *testing.T) {
	client := CreateTestTumblr().Client
	blogAPI := CreateTestAPI(client)
	meta, posts, err := blogAPI.Quote(nil)

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
	id, reblogKey := CreateTestReblogData()

	client := CreateTestTumblr().Client
	blogAPI := CreateTestAPI(client)
	meta, id, err := blogAPI.Reblog(id, reblogKey, "test")

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
