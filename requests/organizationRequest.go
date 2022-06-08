package requests

type OrganizationRequest struct {
	Name     string `json:"name" validate:"required"`
	Period   string `json:"period" validate:"required"`
	Position string `json:"position" validate:"required"`
	Detail   string `json:"detail" validate:"required"`
}
