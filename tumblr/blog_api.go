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

func (blog *BlogApi) send(method string, values *url.Values) ([]byte, error) {
	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + method
	uri += "?" + values.Encode()

	res, err := http.Get(uri)
	if err != nil {
		return make([]byte, 0), err
	}

	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

func (blog *BlogApi) Info() (m *Meta, b *Blog, err error) {
	values := url.Values{}
	values.Add("api_key", blog.client.ConsumerKey)

	method := "/info"
	data, err := blog.send(method, &values)

	var root Root
	json.Unmarshal(data, &root)

	return &root.Meta, &root.Response.Blog, err
}

func (blog *BlogApi) Posts(postType string) (*Meta, *[]Post, error) {
	// https://www.tumblr.com/docs/en/api/v2#posts
	// api.tumblr.com/v2/blog/{base-hostname}/posts[/type]?api_key={key}&[optional-params=]

	values := url.Values{}
	values.Add("api_key", blog.client.ConsumerKey)

	method := "/posts"
	if postType != "" {
		method += "/" + postType
	}

	data, err := blog.send(method, &values)
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
