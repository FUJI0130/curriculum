package userapp

import "time"

type UserRequestUpdate struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Profile  string `json:"Profile"`
}
type SkillRequestUpdate struct {
	TagName    string `json:"TagName"`
	Evaluation uint8  `json:"Evaluation"`
	Years      uint8  `json:"Years"`
}

type CareersRequestUpdate struct {
	Detail string    `json:"Detail"`
	AdFrom time.Time `json:"AdFrom"`
	AdTo   time.Time `json:"AdTo"`
}

type SkillRequestCreate struct {
	TagName    string `json:"TagName"`
	Evaluation uint8  `json:"Evaluation"`
	Years      uint8  `json:"Years"`
}

type CareersRequestCreate struct {
	Detail string    `json:"Detail"`
	AdFrom time.Time `json:"AdFrom"`
	AdTo   time.Time `json:"AdTo"`
}
