package tumblr

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
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

func (blog *BlogApi) get(method string, params *map[string]string) ([]byte, error) {
	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + method

	res, err := blog.client.Get(uri, *params)
	if err != nil {
		return make([]byte, 0), err
	}

	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

func (blog *BlogApi) post(method string, params *map[string]string) ([]byte, error) {
	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + method

	res, err := blog.client.Post(uri, *params)
	if err != nil {
		return make([]byte, 0), err
	}

	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

func (blog *BlogApi) Info() (m *Meta, b *Blog, err error) {
	params := make(map[string]string)
	params["api_key"] = blog.client.ConsumerKey

	method := "/info"
	data, err := blog.get(method, &params)

	var root Root
	json.Unmarshal(data, &root)

	return &root.Meta, &root.Response.Blog, err
}

func (blog *BlogApi) Posts(postType string) (*Meta, *[]Post, error) {
	// https://www.tumblr.com/docs/en/api/v2#posts
	// api.tumblr.com/v2/blog/{base-hostname}/posts[/type]?api_key={key}&[optional-params=]

	params := make(map[string]string)
	params["api_key"] = blog.client.ConsumerKey

	method := "/posts"
	if postType != "" {
		method += "/" + postType
	}

	data, err := blog.get(method, &params)
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

func (blog *BlogApi) Reblog(id int64, reblog_key string, comment string) (*Meta, int64, error) {
	params := make(map[string]string)
	params["id"] = strconv.FormatInt(id, 10)
	params["reblog_key"] = reblog_key

	if comment != "" {
		params["comment"] = comment
	}

	method := "/post/reblog"

	data, err := blog.post(method, &params)
	var response ReblogResponse
	json.Unmarshal(data, &response)

	return &response.Meta, response.Response.Id, err
}
