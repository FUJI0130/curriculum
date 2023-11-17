package userapp

import "time"

type CreateUserRequest struct {
	Name     string                `json:"name"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Skills   []CreateSkillRequest  `json:"skills"`
	Profile  string                `json:"profile"`
	Careers  []CreateCareerRequest `json:"careers"`
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
