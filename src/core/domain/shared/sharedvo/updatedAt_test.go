package sharedvo

import (
	"errors"
	"testing"
	"time"
)

func TestNewUpdatedAt(t *testing.T) {
	// validTime := time.Now().Add(time.Hour)
	// validTime := time.Now().Add(0)
	validTime := time.Now().Add(time.Millisecond)
	tests := []struct {
		name      string
		input     time.Time
		want      UpdatedAt
		wantError error
	}{
		{
			name:      "時間が０",
			input:     time.Time{},
			want:      UpdatedAt(time.Time{}),
			wantError: errors.New("UpdatedAt cannot be zero value"),
		},
		{
			name:      "過去の時間の場合のテスト",
			input:     time.Now().Add(-time.Hour * 24),
			want:      UpdatedAt(time.Time{}),
			wantError: errors.New("UpdatedAt cannot be past date"),
		},
		{
			name:      "有効な時間の場合のテスト",
			input:     validTime,
			want:      (UpdatedAt)(validTime),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdatedAtByVal(tt.input)

			// エラーの有無を確認
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
			name:   "等しい場合のテスト",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now),
			result: true,
		},
		{
			name:   "等しくない場合のテスト",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now.Add(time.Hour)),
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u1.Equal(tt.u2); got != tt.result {
				t.Errorf("expected: %v, got: %v", tt.result, got)
			}
		})
	}
}
