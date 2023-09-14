package userdm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserEmail(t *testing.T) {

	validEmail := UserEmail("test@example.com")
	tests := []struct {
		name          string
		input         string
		expectedError error
		expectedValue *UserEmail
	}{
		{
			name:          "emailが空の場合のテスト",
			input:         "",
			expectedError: errors.New("userEmail cannot be empty"),
			expectedValue: nil,
		},
		{
			name:          "有効なemailか確認するテスト",
			input:         "test@example.com",
			expectedError: nil,
			expectedValue: &validEmail,
		},
		{
			name:          "emailのドメインが有効でない場合のテスト",
			input:         "test@.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			name:          "＠が無い場合のテスト",
			input:         "testexample.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			name:          "emailの長さが256文字以上の場合のテスト",
			input:         "a" + string(make([]rune, 255)) + "@example.com",
			expectedError: errors.New("userEmail length over nameMaxlength"),
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		tt := tt // capture the range variable for parallel execution
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // this allows the subtest to run in parallel
			result, err := NewUserEmail(tt.input)

			assert.Equal(t, tt.expectedError, err)

			if tt.expectedValue != nil {
				assert.Equal(t, tt.expectedValue.String(), result.String())
			}
		})
	}
}
