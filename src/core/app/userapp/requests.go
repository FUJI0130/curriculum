package userapp

import "time"

type updateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}
type updateSkillRequest struct {
	ID         string `json:"id"`
	TagName    string `json:"tagName"`
	Evaluation uint8  `json:"evaluation"`
	Years      uint8  `json:"years"`
}

type updateCareerRequest struct {
	ID     string    `json:"id"`
	Detail string    `json:"detail"`
	AdFrom time.Time `json:"adFrom"`
	AdTo   time.Time `json:"adTo"`
}

type CreateSkillRequest struct {
	TagName    string `json:"tagName"`
	Evaluation uint8  `json:"evaluation"`
	Years      uint8  `json:"years"`
}

type CreateCareerRequest struct {
	Detail string    `json:"detail"`
	AdFrom time.Time `json:"adFrom"`
	AdTo   time.Time `json:"adTo"`
}
