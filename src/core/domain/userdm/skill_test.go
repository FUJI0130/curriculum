package userdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/stretchr/testify/assert"
)

func TestNewSkill(t *testing.T) {
	tests := []struct {
		title      string
		evaluation uint8
		years      uint8
		wantErr    bool
	}{
		{
			title:      "invalid evaluation and valid years",
			evaluation: 0,
			years:      5,
			wantErr:    true,
		},
		{
			title:      "valid evaluation and invalid years",
			evaluation: 3,
			years:      0,
			wantErr:    true,
		},
		{
			title:      "valid evaluation and years",
			evaluation: 3,
			years:      5,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			dummyTagID, err := tagdm.NewTagID()
			assert.NoError(t, err)

			_, err = userdm.NewSkill(dummyTagID, tt.evaluation, tt.years)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
