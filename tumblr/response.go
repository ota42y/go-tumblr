package tumblr

type Root struct{
	Meta Meta
	Response Response
}

type Meta struct{
	Status int
	Msg string
}

type Response struct {
	Blog Blog
}

type Blog struct {
	Title string
	Posts int
	Name string
	Url string
	Updated string // TODO: convert time object
	Description string
	Ask bool
	Ask_anon bool
	Likes int
}
