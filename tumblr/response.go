package tumblr

type Root struct {
	Meta     Meta
	Response Response
}

type Meta struct {
	Status int
	Msg    string
}

type Response struct {
	Blog Blog
	Posts []PostImpl
}

type Blog struct {
	Title       string
	Posts       int
	Name        string
	Url         string
	Updated     int64 // TODO: convert time object
	Description string
	Ask         bool
	Ask_anon    bool
	Likes       int
}

type Post interface {
	Post() string
	Id() int64
	PostUrl() string
	Type() string
	Timestamp() int64
	Date() string
	Format() string
	ReblogKey() string
	Tags() []string
	Bookmarklet() bool
	Mobile() bool
	SourceUrl() string
	SourceTitle() string
	Liked() bool
	State() string
	TotalPosts() int64
}

type PostImpl struct {
	BlogName string
	Id int64
	PostUrl string
	Type string
	Timestamp int64
	Date string
	Format string
	ReblogKey string
	Tags []string
	Bookmarklet bool
	Mobile bool
	SourceUrl string
	SourceTitle string
	Liked bool
	State string
	TotalPosts int64
}

type PostBase struct {
	PostImpl
}
