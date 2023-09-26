package sharedvo

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUpdatedAtByVal(t *testing.T) {
	// validTime := time.Now().Add(time.Hour)
	// validTime := time.Now().Add(0)
	validTime := time.Now().Add(time.Millisecond)
	pastTime := time.Now().Add(-time.Hour * 24)
	tests := []struct {
		title     string
		input     time.Time
		want      UpdatedAt
		wantError error
	}{
		{
			title:     "時間が０",
			input:     time.Time{},
			want:      UpdatedAt(time.Time{}),
			wantError: errors.New("UpdatedAt cannot be zero value"),
		},
		{
			title:     "過去の時間の場合のテスト",
			input:     pastTime,
			want:      UpdatedAt(pastTime),
			wantError: nil,
		},
		{
			title:     "有効な時間の場合のテスト",
			input:     validTime,
			want:      (UpdatedAt)(validTime),
			wantError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			got, err := NewUpdatedAtByVal(tt.input)

			// エラーの有無を確認
			if tt.wantError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.DateTime(), got.DateTime())
			}
		})
	}

}

func TestUpdatedAt_Equal(t *testing.T) {
	now := time.Now()
	tests := []struct {
		title  string
		u1     UpdatedAt
		u2     UpdatedAt
		result bool
	}{
		{
			title:  "等しい場合のテスト",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now),
			result: true,
		},
		{
			title:  "等しくない場合のテスト",
			u1:     UpdatedAt(now),
			u2:     UpdatedAt(now.Add(time.Hour)),
			result: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			got := tt.u1.Equal(tt.u2)
			assert.Equal(t, tt.result, got)
		})
	}
}
