// src/core/domain/tagdm/tag_test.go
package tagdm_test

import (
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTag(t *testing.T) {
	// createdAtとupdatedAtの作成
	createdAt, err := sharedvo.NewCreatedAtByVal(time.Now())
	if err != nil {
		t.Fatalf("createdAtの作成に失敗しました: %v", err)
	}

	startTime := time.Now()
	updatedAt, err := sharedvo.NewUpdatedAtByVal(startTime)
	if err != nil {
		t.Logf("updatedAt.Before(time.Now())にかかった時間: %v", sharedvo.LastDuration)
		t.Fatalf("updatedAtの作成に失敗しました: %v", err)
	}

	tests := []struct {
		name        string
		tagName     string
		createdAt   sharedvo.CreatedAt
		updatedAt   sharedvo.UpdatedAt
		expectError bool
	}{
		{
			name:        "TestTag",
			tagName:     "testTag",
			createdAt:   createdAt,
			updatedAt:   updatedAt,
			expectError: false,
		},
		{
			name:        "タグ名が空のため失敗",
			tagName:     "",
			createdAt:   createdAt,
			updatedAt:   updatedAt,
			expectError: true,
		},
		// ここに他のテストケースを追加することができます。
	}

	for _, test := range tests {
		tag, err := tagdm.NewTag(test.tagName, test.createdAt.DateTime(), test.updatedAt.DateTime())

		if test.expectError {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, test.tagName, tag.Name())
		assert.Equal(t, test.createdAt, tag.CreatedAt())
		assert.Equal(t, test.updatedAt, tag.UpdatedAt())
	}
}
