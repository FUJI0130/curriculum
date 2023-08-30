package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkills(t *testing.T) {
	tests := []struct {
		name       string
		evaluation int
		years      int
		wantErr    bool
	}{
		{
			name:       "invalid evaluation and valid years",
			evaluation: 0,
			years:      5,
			wantErr:    true,
		},
		{
			name:       "valid evaluation and invalid years",
			evaluation: 3,
			years:      0,
			wantErr:    true,
		},
		{
			name:       "valid evaluation and years",
			evaluation: 3,
			years:      5,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userdm.NewSkills(tt.evaluation, tt.years)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
