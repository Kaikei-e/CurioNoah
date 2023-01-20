package request

type SingleFeed struct {
	URL    string `json:"url" ,validate:"required,url"`
	UserID string `json:"user_id"`
}
