package link

type LinkCreationRequest struct {
	Url string `json:"url" validate:"required"`
}
