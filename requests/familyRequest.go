package requests

type FamilyRequest struct {
	Status     string `json:"status" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Birthplace string `json:"birthplace" validate:"required"`
	Birthdate  string `json:"birthdate" validate:"required"`
	Education  string `json:"education" validate:"required"`
	Job        string `json:"job"`
}
