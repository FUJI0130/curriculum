package mentorapp

import "time"

type CreateMentorRecruitmentRequest struct {
	Title                 string    `json:"title"`
	CategoryID            string    `json:"category_id"`
	BudgetFrom            uint32    `json:"budget_from"`
	BudgetTo              uint32    `json:"budget_to"`
	ApplicationPeriodFrom time.Time `json:"application_period_from"`
	ApplicationPeriodTo   time.Time `json:"application_period_to"`
	ConsultationFormat    uint8     `json:"consultation_format"`
	ConsultationMethod    uint8     `json:"consultation_method"`
	Description           string    `json:"description"`
	Status                uint8     `json:"status"`
	TagIDs                []string  `json:"tag_ids"`
}
