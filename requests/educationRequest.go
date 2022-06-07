package requests

type EducationRequest struct {
	Grade    string `json:"grade" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Province string `json:"province" validate:"required"`
	City     string `json:"city" validate:"required"`
	Enter    string `json:"enter" validate:"required"`
	Graduate string `json:"graduate" validate:"required"`
}
