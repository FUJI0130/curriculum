package userapp

import "time"

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}
type UpdateSkillRequest struct {
	ID         string `json:"id"`
	TagID      string `json:"tagId"`
	Evaluation uint8  `json:"evaluation"`
	Years      uint8  `json:"years"`
}

type UpdateCareerRequest struct {
	ID     string    `json:"id"`
	Detail string    `json:"detail"`
	AdFrom time.Time `json:"adFrom"`
	AdTo   time.Time `json:"adTo"`
}

type UpdateUserRequestData struct {
	UpdateData struct {
		Users   UpdateUserRequest     `json:"users"`
		Skills  []UpdateSkillRequest  `json:"skills"`
		Careers []UpdateCareerRequest `json:"careers"`
	} `json:"updateData"`
}
