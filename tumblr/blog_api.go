package tumblr

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"strconv"
)

// BlogAPI is tumblr blog's api
type BlogAPI struct {
	Host   string
	client *Client
	token  *oauth.AccessToken
}

// NewblogAPI return blog's api
func NewblogAPI(host string, client *Client, token *oauth.AccessToken) *BlogAPI {
	return &BlogAPI{
		Host:   host,
		client: client,
		token:  token,
	}
}

func (blog *BlogAPI) get(method string, params *map[string]string) ([]byte, error) {
	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + method

	res, err := blog.client.Get(uri, *params, blog.token)
	if err != nil {
		return make([]byte, 0), err
	}

	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

func (blog *BlogAPI) post(method string, params *map[string]string) ([]byte, error) {
	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + method

	res, err := blog.client.Post(uri, *params, blog.token)
	if err != nil {
		return make([]byte, 0), err
	}

	data, err := ioutil.ReadAll(res.Body)
	return data, err
}

// Info return blog's infomation
func (blog *BlogAPI) Info() (m *Meta, b *Blog, err error) {
	params := make(map[string]string)
	params["api_key"] = blog.client.GetConsumerKey()

	method := "/info"
	data, err := blog.get(method, &params)

	var root Root
	json.Unmarshal(data, &root)

	return &root.Meta, &root.Response.Blog, err
}

// Posts return blog's posts (https://www.tumblr.com/docs/en/api/v2#posts)
func (blog *BlogAPI) Posts(postType string, params *map[string]string) (*Meta, *[]Post, error) {
	// https://www.tumblr.com/docs/en/api/v2#posts
	// api.tumblr.com/v2/blog/{base-hostname}/posts[/type]?api_key={key}&[optional-params=]

	if params == nil {
		p := make(map[string]string)
		params = &p
	}
	(*params)["api_key"] = blog.client.GetConsumerKey()

	method := "/posts"
	if postType != "" {
		method += "/" + postType
	}

	data, err := blog.get(method, params)
	var root Root
	json.Unmarshal(data, &root)

	return &root.Meta, &root.Response.Posts, err
}

// Photo return photo api response
func (blog *BlogAPI) Photo(params *map[string]string) (*Meta, *[]Post, error) {
	return blog.Posts("photo", params)
}

// Quote is quote specific post
func (blog *BlogAPI) Quote(params *map[string]string) (*Meta, *[]Post, error) {
	return blog.Posts("quote", params)
}

// Reblog is reblog specific post
func (blog *BlogAPI) Reblog(id int64, reblogKey string, comment string) (*Meta, int64, error) {
	params := make(map[string]string)
	params["id"] = strconv.FormatInt(id, 10)
	params["reblog_key"] = reblogKey

	if comment != "" {
		params["comment"] = comment
	}

	method := "/post/reblog"

	data, err := blog.post(method, &params)
	var response ReblogResponse
	json.Unmarshal(data, &response)

	return &response.Meta, response.Response.ID, err
}
