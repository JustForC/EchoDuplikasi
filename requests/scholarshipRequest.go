package requests

type ScholarshipRequest struct {
	Name            string `json:"name" validate:"required"`
	StartStepOne    string `json:"startStepOne" validate:"required"`
	StartStepTwo    string `json:"startStepTwo" validate:"required"`
	EndStepOne      string `json:"endStepOne" validate:"required"`
	EndStepTwo      string `json:"endStepTwo" validate:"required"`
	AnnounceStepOne string `json:"announceStepOne" validate:"required"`
	AnnounceStepTwo string `json:"announceStepTwo" validate:"required"`
	OnlineTest      string `json:"onlineTest" validate:"required"`
}
