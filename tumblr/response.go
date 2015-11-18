package tumblr

// Root is response root
type Root struct {
	Meta     Meta
	Response Response
}

// Meta is meta data struct
type Meta struct {
	Status int
	Msg    string
}

// Response is tumblr response struct
type Response struct {
	Blog  Blog
	Posts []Post
}

// Blog is tumbler blog struct
type Blog struct {
	Title       string
	Posts       int
	Name        string
	URL         string
	Updated     int64 // TODO: convert time object
	Description string
	Ask         bool
	AskAnon     bool `json:"ask_anon"`
	Likes       int
}

// Post is tumblr post struct
type Post struct {
	BlogName    string
	ID          int64
	PostURL     string `json:"post_url"`
	Type        string
	Timestamp   int64
	Date        string
	Format      string
	ReblogKey   string `json:"reblog_key"`
	Tags        []string
	Bookmarklet bool
	Mobile      bool
	SourceURL   string
	SourceTitle string
	Liked       bool
	State       string
	TotalPosts  int64

	Caption        string
	ImagePermalink string `json:"image_permalink"`

	// Post
	Title string
	Body  string

	// Photo
	Photos []Photo `json:"photos"`

	// Quote
	Text   string
	Source string
}

// Photo is post's photo struct
type Photo struct {
	Caption  string    `json:"caption"`
	AltSizes []AltSize `json:"alt_sizes"`
}

// AltSize is photo's alt_sizez struct
type AltSize struct {
	Width  int
	Height int
	URL    string
}

// ReblogResponse is reblog api response
type ReblogResponse struct {
	Meta     Meta
	Response struct {
		ID int64
	}
}
