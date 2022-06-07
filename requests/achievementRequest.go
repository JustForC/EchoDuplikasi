package requests

type AchievementRequest struct {
	Name      string `json:"name" validate:"required"`
	Organizer string `json:"organizer" validate:"required"`
	Level     string `json:"level" validate:"required"`
}
