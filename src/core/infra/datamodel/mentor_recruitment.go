package datamodel

import "time"

type MentorRecruitment struct {
	ID                    string    `db:"id"`
	Title                 string    `db:"name"`
	CategoryID            string    `db:"tag_id"`
	BudgetFrom            uint32    `db:"budget_from"`
	BudgetTo              uint32    `db:"budget_to"`
	ApplicationPeriodFrom time.Time `db:"application_period_from"`
	ApplicationPeriodTo   time.Time `db:"application_period_to"`
	ConsultationFormat    uint8     `db:"consultation_format"`
	ConsultationMethod    uint8     `db:"consultation_method"`
	Description           string    `db:"description"`
	Status                uint8     `db:"status"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}
