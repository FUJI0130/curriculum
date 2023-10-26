package userapp

import "time"

type SkillRequest struct {
	TagName    string
	Evaluation uint8
	Years      uint8
}

type CareersRequest struct {
	Detail string
	AdFrom time.Time
	AdTo   time.Time
}
