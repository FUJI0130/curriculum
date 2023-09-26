package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillYear(t *testing.T) {
	tests := []struct {
		title   string
		years   uint8
		wantErr bool
	}{
		{
			title:   "zero years",
			years:   0,
			wantErr: true,
		},
		{
			title:   "over 100 years",
			years:   101,
			wantErr: true,
		},
		{
			title:   "valid years",
			years:   5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture the range variable for parallel execution
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
			_, err := userdm.NewSkillYear(tt.years)

			// Check if the error is nil or not
			assert.Equal(t, tt.wantErr, err != nil)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSkillYear_Equal(t *testing.T) {
	years1, err1 := userdm.NewSkillYear(3)
	if err1 != nil {
		t.Fatalf("Failed to create SkillYear for years1: %v", err1)
	}

	years2, err2 := userdm.NewSkillYear(3)
	if err2 != nil {
		t.Fatalf("Failed to create SkillYear for years2: %v", err2)
	}

	years3, err3 := userdm.NewSkillYear(4)
	if err3 != nil {
		t.Fatalf("Failed to create SkillYear for years3: %v", err3)
	}

	assert.True(t, years1.Equal(&years2))
	assert.False(t, years1.Equal(&years3))
	assert.False(t, years1.Equal(nil))
	assert.True(t, (*userdm.SkillYear)(nil).Equal(nil))
}
