package requests

type TalentRequest struct {
	Name string `json:"name" validate:"required"`
}
