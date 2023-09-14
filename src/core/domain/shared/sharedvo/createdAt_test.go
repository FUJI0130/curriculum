package sharedvo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCreatedAt(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		input     time.Time
		expectErr bool
		errorMsg  string
	}{
		{
			name:      "valid date",
			input:     now.Add(-1 * time.Hour),
			expectErr: false,
		},
		{
			name:      "zero date",
			input:     time.Time{},
			expectErr: true,
			errorMsg:  "CreatedAt cannot be zero value",
		},
		{
			name:      "future date",
			input:     now.Add(1 * time.Hour),
			expectErr: true,
			errorMsg:  "CreatedAt cannot be future date",
		},
	}

	for _, tt := range tests {
		tt := tt // capture the range variable for parallel execution
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
			_, err := NewCreatedAtByVal(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreatedAt_Equal(t *testing.T) {
	date1, _ := NewCreatedAtByVal(time.Now())
	date2, _ := NewCreatedAtByVal(time.Now())
	date3, _ := NewCreatedAtByVal(time.Now().Add(-1 * time.Hour))

	tests := []struct {
		name   string
		date1  CreatedAt
		date2  CreatedAt
		result bool
	}{
		{
			name:   "日付が等しい",
			date1:  date1,
			date2:  date2,
			result: false,
		},
		{
			name:   "異なる日付",
			date1:  date1,
			date2:  date3,
			result: false,
		},
		{
			name:   "両方データが空だった場合",
			date1:  CreatedAt(time.Time{}),
			date2:  CreatedAt(time.Time{}),
			result: true,
		},
		{
			name:   "date2が空データ",
			date1:  date1,
			date2:  CreatedAt(time.Time{}),
			result: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture the range variable for parallel execution
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
			assert.Equal(t, tt.result, tt.date1.Equal(tt.date2))
		})
	}
}
