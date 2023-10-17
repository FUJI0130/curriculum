package userdm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserEmail(t *testing.T) {

	validEmail := UserEmail("test@example.com")
	tests := []struct {
		title         string
		input         string
		expectedError error
		expectedValue *UserEmail
	}{
		{
			title:         "emailが空の場合のテスト",
			input:         "",
			expectedError: errors.New("userEmail cannot be empty"),
			expectedValue: nil,
		},
		{
			title:         "有効なemailか確認するテスト",
			input:         "test@example.com",
			expectedError: nil,
			expectedValue: &validEmail,
		},
		{
			title:         "emailのドメインが有効でない場合のテスト",
			input:         "test@.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			title:         "＠が無い場合のテスト",
			input:         "testexample.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			title:         "emailの長さが256文字以上の場合のテスト",
			input:         "a" + string(make([]rune, 255)) + "@example.com",
			expectedError: errors.New("userEmail length over nameMaxlength"),
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			result, err := NewUserEmail(tt.input)

			if tt.expectedError == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedError, err)
			}

			if tt.expectedValue != nil {
				assert.Equal(t, tt.expectedValue.String(), result.String())
			}
		})
	}
}
