package apiexcahnge

type FeedRegister struct {
	URL    string `json:"url" ,validate:"required,url"`
	UserID string `json:"user_id"`
}

type FeedResponse struct {
	IsSuccess bool `json:"is_success"`
	// TODO think about the domain
}
