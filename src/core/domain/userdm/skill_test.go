package userdm_test

import (
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkill(t *testing.T) {
	tests := []struct {
		name       string
		evaluation uint8
		years      uint8
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
			dummyTagID, _ := tagdm.NewTagID()
			dummyUserID, _ := userdm.NewUserID()
			dummyCreatedAt := time.Now()
			dummyUpdatedAt := time.Now()

			_, err := userdm.NewSkill(dummyTagID, dummyUserID, tt.evaluation, tt.years, dummyCreatedAt, dummyUpdatedAt)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
