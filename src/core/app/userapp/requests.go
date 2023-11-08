package userapp

import "time"

type UpdateUserRequest struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Profile  string `json:"Profile"`
}
type UpdateSkillRequest struct {
	ID         string `json:"ID"`
	TagName    string `json:"TagName"`
	Evaluation uint8  `json:"Evaluation"`
	Years      uint8  `json:"Years"`
}

type UpdateCareerRequest struct {
	ID     string    `json:"ID"`
	Detail string    `json:"Detail"`
	AdFrom time.Time `json:"AdFrom"`
	AdTo   time.Time `json:"AdTo"`
}

type CreateSkillRequest struct {
	TagName    string `json:"TagName"`
	Evaluation uint8  `json:"Evaluation"`
	Years      uint8  `json:"Years"`
}

type CreateCareerRequest struct {
	Detail string    `json:"Detail"`
	AdFrom time.Time `json:"AdFrom"`
	AdTo   time.Time `json:"AdTo"`
}
