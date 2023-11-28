package datamodel

import "time"

type MentorRecruitment struct {
	ID                    string    `db:"id"`
	Title                 string    `db:"title"`
	CategoryID            string    `db:"category_id"`
	BudgetFrom            int       `db:"budget_from"`
	BudgetTo              int       `db:"budget_to"`
	ApplicationPeriodFrom time.Time `db:"application_period_from"`
	ApplicationPeriodTo   time.Time `db:"application_period_to"`
	ConsultationFormat    int       `db:"consultation_format"`
	ConsultationMethod    int       `db:"consultation_method"`
	Description           string    `db:"description"`
	Status                int       `db:"status"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}
