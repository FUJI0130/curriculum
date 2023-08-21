package sharedvo

import (
	"testing"
	"time"
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
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCreatedAt(tt.input)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, but got none")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Did not expect an error, but got: %v", err)
			} else if tt.expectErr && err != nil && err.Error() != tt.errorMsg {
				t.Errorf("Expected error message %q, but got %q", tt.errorMsg, err.Error())
			}
		})
	}
}

func TestCreatedAt_Equal(t *testing.T) {
	date1, _ := NewCreatedAt(time.Now())
	date2, _ := NewCreatedAt(time.Now())
	date3, _ := NewCreatedAt(time.Now().Add(-1 * time.Hour))

	tests := []struct {
		name   string
		date1  *CreatedAt
		date2  *CreatedAt
		result bool
	}{
		{
			name:   "equal dates",
			date1:  date1,
			date2:  date2,
			result: false,
		},
		{
			name:   "different dates",
			date1:  date1,
			date2:  date3,
			result: false,
		},
		{
			name:   "both nil",
			date1:  nil,
			date2:  nil,
			result: true,
		},
		{
			name:   "one nil",
			date1:  date1,
			date2:  nil,
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.date1.Equal(tt.date2) != tt.result {
				t.Errorf("Expected %v, but got %v", tt.result, !tt.result)
			}
		})
	}
}
