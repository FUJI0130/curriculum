package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillsEvaluation(t *testing.T) {
	tests := []struct {
		name    string
		value   uint8
		wantErr bool
	}{
		{
			name:    "below minimum",
			value:   0,
			wantErr: true,
		},
		{
			name:    "above maximum",
			value:   6,
			wantErr: true,
		},
		{
			name:    "valid evaluation",
			value:   3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userdm.NewSkillsEvaluation(tt.value)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSkillsEvaluation_Equal(t *testing.T) {
	evaluation1, _ := userdm.NewSkillsEvaluation(3)
	evaluation2, _ := userdm.NewSkillsEvaluation(3)
	evaluation3, _ := userdm.NewSkillsEvaluation(4)

	assert.True(t, evaluation1.Equal(evaluation2))
	assert.False(t, evaluation1.Equal(evaluation3))
	assert.False(t, evaluation1.Equal(nil))
	assert.True(t, (*userdm.SkillsEvaluation)(nil).Equal(nil))
}
