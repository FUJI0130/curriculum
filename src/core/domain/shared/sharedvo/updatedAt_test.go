package sharedvo

import (
	"errors"
	"testing"
	"time"
)

func TestNewUpdatedAt(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		want      *UpdatedAt
		wantError error
	}{
		{
			name:      "zero time",
			input:     time.Time{},
			want:      nil,
			wantError: errors.New("UpdatedAt cannot be zero value"),
		},
		{
			name:      "future time",
			input:     time.Now().Add(time.Hour * 24),
			want:      nil,
			wantError: errors.New("UpdatedAt cannot be future date"),
		},
		{
			name:      "valid time",
			input:     time.Now().Add(-time.Hour),
			want:      NewUpdatedAt(time.Now().Add(-time.Hour)),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdatedAt(tt.input)
			if tt.wantError != nil {
				if err == nil || tt.wantError.Error() != err.Error() {
					t.Errorf("expected error: %v, got: %v", tt.wantError, err)
				}
				return
			}

			if !got.DateTime().Equal(tt.want.DateTime()) {
				t.Errorf("expected: %v, got: %v", tt.want.DateTime(), got.DateTime())
			}
		})
	}
}

func TestUpdatedAt_Equal(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		u1     UpdatedAt
		u2     UpdatedAt
		result bool
	}{
		{
			name:   "equal",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now),
			result: true,
		},
		{
			name:   "not equal",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now.Add(time.Hour)),
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u1.Equal(&tt.u2); got != tt.result {
				t.Errorf("expected: %v, got: %v", tt.result, got)
			}
		})
	}
}
