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
		want      *UpdatedAt
		wantError error
	}{
		{
			name:      "時間が０",
			input:     time.Time{},
			want:      nil,
			wantError: errors.New("UpdatedAt cannot be zero value"),
		},
		{
			name:      "過去の時間の場合のテスト",
			input:     time.Now().Add(-time.Hour * 24),
			want:      nil,
			wantError: errors.New("UpdatedAt cannot be past date"),
		},
		{
			name:      "有効な時間の場合のテスト",
			input:     validTime,
			want:      (*UpdatedAt)(&validTime),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdatedAt(tt.input)

			// エラーの有無を確認
			if tt.wantError != nil {
				if err == nil || tt.wantError.Error() != err.Error() {
					t.Errorf("expected error: %v, got: %v", tt.wantError, err)
				}
				return
			}

			// got または tt.want が nil の場合の処理
			if got == nil || tt.want == nil {
				if (got != nil && tt.want == nil) || (got == nil && tt.want != nil) {
					t.Errorf("expected: %v, got: %v", tt.want, got)
				}
				return
			}

			// DateTimeメソッドとEqualメソッドを使用した比較
			if !got.DateTime().Equal(tt.want.DateTime()) {
				t.Errorf("expected: %v, got: %v", tt.want.DateTime(), got.DateTime())
			}
		})
	}

}
func TestUpdatedAt_SetDateTime(t *testing.T) {
	tests := []struct {
		name      string
		original  time.Time
		newTime   time.Time
		wantError error
	}{
		{
			name:      "時間が０",
			original:  time.Now(),
			newTime:   time.Time{},
			wantError: errors.New("UpdatedAt cannot be zero value"),
		},
		{
			name:      "過去の時間の場合のテスト",
			original:  time.Now(),
			newTime:   time.Now().Add(-time.Hour),
			wantError: errors.New("UpdatedAt cannot be past date"),
		},
		{
			name:      "有効な時間の場合のテスト",
			original:  time.Now(),
			newTime:   time.Now().Add(time.Hour),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UpdatedAt(tt.original)
			err := u.SetDateTime(tt.newTime)

			if tt.wantError != nil {
				if err == nil || tt.wantError.Error() != err.Error() {
					t.Errorf("expected error: %v, got: %v", tt.wantError, err)
				}
				return
			}

			if !u.DateTime().Equal(tt.newTime) {
				t.Errorf("expected: %v, got: %v", tt.newTime, u.DateTime())
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
			if got := tt.u1.Equal(&tt.u2); got != tt.result {
				t.Errorf("expected: %v, got: %v", tt.result, got)
			}
		})
	}
}
