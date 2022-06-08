package requests

type TrainingRequest struct {
	Name        string `json:"name" validate:"required"`
	Period      string `json:"period" validate:"required"`
	Year        string `json:"year" validate:"required"`
	Organizer   string `json:"organizer" validate:"required"`
	Certificate string `json:"certificate" validate:"required"`
}
