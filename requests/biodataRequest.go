package requests

type BiodataRequest struct {
	Name             string `json:"name" validate:"required"`
	Nickname         string `json:"nickname" validate:"required"`
	Gender           string `json:"gender" validate:"required"`
	Birthplace       string `json:"birthplace" validate:"required"`
	Birthdate        string `json:"birthdate" validate:"required"`
	Telephone        string `json:"telephone" validate:"required,min=10,max=13"`
	Email            string `json:"email" validate:"required,email"`
	Id_Type          string `json:"id_type" validate:"required"`
	Id_Number        string `json:"id_number" validate:"required"`
	Address          string `json:"address" validate:"required"`
	Post_Code        string `json:"post_code" validate:"required"`
	District         string `json:"district" validate:"required"`
	City             string `json:"city" validate:"required"`
	Province         string `json:"province" validate:"required"`
	Living_Address   string `json:"living_address" validate:"required"`
	Living_Post_Code string `json:"living_post_code" validate:"required"`
	Living_District  string `json:"living_district" validate:"required"`
	Living_City      string `json:"living_city" validate:"required"`
	Living_Province  string `json:"living_province" validate:"required"`
	Entrance         string `json:"entrance" validate:"required"`
	Entrance_Number  string `json:"entrance_number" validate:"required"`
	Major            string `json:"major" validate:"required"`
	University       string `json:"university" validate:"required"`
	Same_Address     bool   `json:"same_address"`
}
