package requests

type NetworthRequest struct {
	Value int64 `json:"value" validate:"required"`
}
