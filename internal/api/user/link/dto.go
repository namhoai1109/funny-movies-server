package link

type LinkCreationRequest struct {
	Url string `json:"url" validate:"required"`
}

type YtbOembedResponse struct {
	Title string `json:"title"`
}
