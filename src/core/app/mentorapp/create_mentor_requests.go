package mentorapp

import "time"

type CreateMentorRecruitmentRequest struct {
	Title                 string    `json:"title"`
	CategoryName          string    `json:"category_name"`
	BudgetFrom            int       `json:"budget_from"`
	BudgetTo              int       `json:"budget_to"`
	ApplicationPeriodFrom time.Time `json:"application_period_from"`
	ApplicationPeriodTo   time.Time `json:"application_period_to"`
	ConsultationFormat    int       `json:"consultation_format"`
	ConsultationMethod    int       `json:"consultation_method"`
	Description           string    `json:"description"`
	Status                int       `json:"status"`
	TagNames              []string  `json:"tag_names"`
}
