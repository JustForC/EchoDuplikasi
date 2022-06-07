package requests

type LanguageRequest struct {
	Name   string `json:"name" validate:"required"`
	Talk   string `json:"talk" validate:"required"`
	Write  string `json:"write" validate:"required"`
	Read   string `json:"read" validate:"required"`
	Listen string `json:"listen" validate:"required"`
}
