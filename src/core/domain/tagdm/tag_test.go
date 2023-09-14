package tagdm_test

import (
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTag(t *testing.T) {
	tests := []struct {
		name        string
		tagName     string
		expectError bool
	}{
		{
			name:        "正しいタグ名",
			tagName:     "testTag",
			expectError: false,
		},
		{
			name:        "タグ名が空",
			tagName:     "",
			expectError: true,
		},
	}

	for _, test := range tests {
		test := test // capture the range variable for parallel execution
		t.Run(test.name, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
			tag, err := tagdm.NewTag(test.tagName)

			if test.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, test.tagName, tag.Name())
			assert.False(t, tag.CreatedAt().DateTime().IsZero()) // CreatedAtが設定されていてゼロではないことを確認
			assert.False(t, tag.UpdatedAt().DateTime().IsZero()) // UpdatedAtが設定されていてゼロではないことを確認
		})
	}
}
