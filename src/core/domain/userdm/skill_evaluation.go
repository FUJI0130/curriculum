// skills_evaluation.go

package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

const (
	MinSkillEvaluationValue = 1
	MaxSkillEvaluationValue = 5
)

type SkillEvaluation uint8

func NewSkillEvaluation(value uint8) (SkillEvaluation, error) {
	if value < MinSkillEvaluationValue || value > MaxSkillEvaluationValue {
		return 0, customerrors.NewUnprocessableEntityErrorf("[NewSkillEvaluation] SkillEvaluation is out of range value is : %d", int(value))
	}
	return SkillEvaluation(value), nil
}

func (evaluation SkillEvaluation) Value() uint8 {
	return uint8(evaluation)
}

func (evaluation1 *SkillEvaluation) Equal(evaluation2 *SkillEvaluation) bool {
	if evaluation1 == nil && evaluation2 == nil {
		return true
	}

	if evaluation1 == nil || evaluation2 == nil {
		return false
	}
	return *evaluation1 == *evaluation2
}
