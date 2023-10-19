package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillEvaluation(t *testing.T) {
	tests := []struct {
		title      string
		evaluation uint8
		wantErr    bool
	}{
		{
			title:      "below minimum",
			evaluation: 0,
			wantErr:    true,
		},
		{
			title:      "above maximum",
			evaluation: 6,
			wantErr:    true,
		},
		{
			title:      "valid evaluation",
			evaluation: 3,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			_, err := userdm.NewSkillEvaluation(tt.evaluation)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSkillEvaluation_Equal(t *testing.T) {
	evaluation1, _ := userdm.NewSkillEvaluation(3)
	evaluation2, _ := userdm.NewSkillEvaluation(3)
	evaluation3, _ := userdm.NewSkillEvaluation(4)

	assert.True(t, evaluation1.Equal(&evaluation2))
	assert.False(t, evaluation1.Equal(&evaluation3))
	assert.False(t, evaluation1.Equal(nil))
	assert.True(t, (*userdm.SkillEvaluation)(nil).Equal(nil))
}
