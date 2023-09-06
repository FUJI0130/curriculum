package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkillYears(t *testing.T) {
	tests := []struct {
		name    string
		years   uint8
		wantErr bool
	}{
		{
			name:    "zero years",
			years:   0,
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
			_, err := userdm.NewSkillYears(tt.years)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSkillYears_Equal(t *testing.T) {
	years1, _ := userdm.NewSkillYears(3)
	years2, _ := userdm.NewSkillYears(3)
	years3, _ := userdm.NewSkillYears(4)

	assert.True(t, years1.Equal(years2))
	assert.False(t, years1.Equal(years3))
	assert.False(t, years1.Equal(nil))
	assert.True(t, (*userdm.SkillYears)(nil).Equal(nil))
}
