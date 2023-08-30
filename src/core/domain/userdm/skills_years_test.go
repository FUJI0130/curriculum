package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillsYears(t *testing.T) {
	tests := []struct {
		name    string
		years   int
		wantErr bool
	}{
		{
			name:    "zero years",
			years:   0,
			wantErr: true,
		},
		{
			name:    "negative years",
			years:   -1,
			wantErr: true,
		},
		{
			name:    "over 100 years",
			years:   101,
			wantErr: true,
		},
		{
			name:    "valid years",
			years:   5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userdm.NewSkillsYears(tt.years)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSkillsYears_Equal(t *testing.T) {
	years1, _ := userdm.NewSkillsYears(3)
	years2, _ := userdm.NewSkillsYears(3)
	years3, _ := userdm.NewSkillsYears(4)

	assert.True(t, years1.Equal(years2))
	assert.False(t, years1.Equal(years3))
	assert.False(t, years1.Equal(nil))
	assert.True(t, (*userdm.SkillsYears)(nil).Equal(nil))
}
