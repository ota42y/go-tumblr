package tumblr

type Blog struct{
	Host string
}

func NewBlog(host string) (*Blog){
	return &Blog{
		Host: host,
	}
}

func (blog *Blog) Info() (res *Response){
	return &Response{}
}
