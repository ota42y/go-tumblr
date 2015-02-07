package tumblr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type BlogApi struct {
	Host   string
	client *Client
}

func NewBlogApi(host string, client *Client) *BlogApi {
	return &BlogApi{
		Host:   host,
		client: client,
	}
}

func (blog *BlogApi) Info() (m *Meta, b *Blog, err error) {
	values := url.Values{}
	values.Add("api_key", blog.client.ConsumerKey)

	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + "/info?" + values.Encode()

	res, err := http.Get(uri)

	var r Root
	dec := json.NewDecoder(res.Body)
	dec.Decode(&r)

	return &r.Meta, &r.Response.Blog, err
}

func (blog *BlogApi) Posts(postType string) (*Meta, *[]Post, error) {
	// https://www.tumblr.com/docs/en/api/v2#posts
	// api.tumblr.com/v2/blog/{base-hostname}/posts[/type]?api_key={key}&[optional-params=]

	values := url.Values{}
	values.Add("api_key", blog.client.ConsumerKey)

	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + "/posts"
	if postType != "" {
		uri += "/" + postType
	}
	uri += "?" + values.Encode()

	res, err := http.Get(uri)
	if err != nil {
		return nil, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	var root Root
	json.Unmarshal(data, &root)

	return &root.Meta, &root.Response.Posts, err
}

func (blog *BlogApi) Photo() (*Meta, *[]Post, error) {
	return blog.Posts("photo")
}

func (blog *BlogApi) Quote() (*Meta, *[]Post, error) {
	return blog.Posts("quote")
}
