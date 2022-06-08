package requests

type SocialMediaRequest struct {
	Instagram string `json:"instagram" validate:"required"`
	Facebook  string `json:"facebook" validate:"required"`
	Tiktok    string `json:"tiktok" validate:"required"`
	Twitter   string `json:"twitter" validate:"required"`
}
