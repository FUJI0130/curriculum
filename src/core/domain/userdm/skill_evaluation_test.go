package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillEvaluation(t *testing.T) {
	tests := []struct {
		name       string
		evaluation uint8
		wantErr    bool
	}{
		{
			name:       "below minimum",
			evaluation: 0,
			wantErr:    true,
		},
		{
			name:       "above maximum",
			evaluation: 6,
			wantErr:    true,
		},
		{
			name:       "valid evaluation",
			evaluation: 3,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture the range variable for parallel execution
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
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
