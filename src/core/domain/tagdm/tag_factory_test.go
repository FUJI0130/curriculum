package tagdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/stretchr/testify/assert"
)

func TestGenWhenCreateTag(t *testing.T) {
	tests := []struct {
		name    string
		tagName string
		wantErr bool
	}{
		{
			name:    "正常なタグの生成",
			tagName: "TestTag",
			wantErr: false,
		},
		{
			name:    "空のタグ名",
			tagName: "",
			wantErr: true,
		},
		{
			name:    "タグ名が長すぎる",
			tagName: "ThisTagNameIsDefinitelyTooLongForTheSystemToHandleProperly",
			wantErr: true,
		},
		// 他のエラーケースやバリデーションケースを追加できます
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := tagdm.GenWhenCreateTag(tt.tagName)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTestNewTag(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		tagName string
		wantErr bool
	}{
		{
			name:    "正常なテストタグの生成",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			tagName: "TestTag",
			wantErr: false,
		},
		{
			name:    "空のタグ名",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			tagName: "",
			wantErr: true,
		},
		{
			name:    "タグ名が長すぎる",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			tagName: "ThisTagNameIsDefinitelyTooLongForTheSystemToHandleProperly",
			wantErr: true,
		},
		// 他のエラーケースやバリデーションケースを追加できます
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := tagdm.TestNewTag(tt.id, tt.tagName)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
