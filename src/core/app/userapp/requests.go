package userapp

import "time"

type UserRequest struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Profile  string `json:"Profile"`
}

type SkillRequest struct {
	TagName    string `json:"TagName"`
	Evaluation uint8  `json:"Evaluation"`
	Years      uint8  `json:"Years"`
}

type CareersRequest struct {
	Detail string    `json:"Detail"`
	AdFrom time.Time `json:"AdFrom"`
	AdTo   time.Time `json:"AdTo"`
}
