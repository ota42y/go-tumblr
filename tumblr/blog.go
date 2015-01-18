package tumblr


import (
	"net/http"
	"net/url"
)

type Blog struct{
	Host string
	client *Client
}

func NewBlog(host string, client *Client) (*Blog){
	return &Blog{
		Host: host,
		client: client,
	}
}

func (blog *Blog) Info() (res *http.Response, err error){
	values := url.Values{}
	values.Add("api_key", blog.client.ConsumerKey)

	uri := "http://api.tumblr.com/v2/blog/" + blog.Host + "/info?" + values.Encode()

	res, err = http.Get(uri)
	return
}
