package tumblr

import (
	"encoding/json"
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
