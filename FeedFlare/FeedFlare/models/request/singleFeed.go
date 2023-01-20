package request

type SingleFeed struct {
	URL string `json:"url" ,validate:"required,url"`
}
