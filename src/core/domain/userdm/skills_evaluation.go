// skills_evaluation.go

package userdm

import (
	"errors"
	"fmt"
)

type SkillsEvaluation int

// スキルの評価は1から10の間とします。
const (
	MinSkillsEvaluationValue = 1
	MaxSkillsEvaluationValue = 5
)

func NewSkillsEvaluation(value int) (*SkillsEvaluation, error) {
	if value < MinSkillsEvaluationValue || value > MaxSkillsEvaluationValue {
		return nil, errors.New(fmt.Sprintf("SkillsEvaluation must be between %d and %d", MinSkillsEvaluationValue, MaxSkillsEvaluationValue))
	}

	evaluation := SkillsEvaluation(value)
	return &evaluation, nil
}

func (evaluation SkillsEvaluation) Value() int {
	return int(evaluation)
}

func (evaluation1 *SkillsEvaluation) Equal(evaluation2 *SkillsEvaluation) bool {
	// 両方のポインタがnilの場合はtrueを返す
	if evaluation1 == nil && evaluation2 == nil {
		return true
	}

	// 片方のポインタだけがnilの場合はfalseを返す
	if evaluation1 == nil || evaluation2 == nil {
		return false
	}
	return *evaluation1 == *evaluation2
}
