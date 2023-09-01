// skills_evaluation.go

package userdm

import (
	"errors"
	"fmt"
)

type SkillEvaluation uint8

// スキルの評価は1から10の間とします。
const (
	MinSkillEvaluationValue = 1
	MaxSkillEvaluationValue = 5
)

func NewSkillEvaluation(value uint8) (*SkillEvaluation, error) {
	if value < MinSkillEvaluationValue || value > MaxSkillEvaluationValue {
		return nil, errors.New(fmt.Sprintf("SkillEvaluation must be between %d and %d", MinSkillEvaluationValue, MaxSkillEvaluationValue))
	}

	evaluation := SkillEvaluation(value)
	return &evaluation, nil
}

func (evaluation SkillEvaluation) Value() uint8 {
	return uint8(evaluation)
}

func (evaluation1 *SkillEvaluation) Equal(evaluation2 *SkillEvaluation) bool {
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
