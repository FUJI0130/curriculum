package userdm

import (
	"errors"
	"testing"
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
			name:          "empty email",
			input:         "",
			expectedError: errors.New("userEmail cannot be empty"),
			expectedValue: nil,
		},
		{
			name:          "valid email",
			input:         "test@example.com",
			expectedError: nil,
			expectedValue: &validEmail,
		},
		{
			name:          "invalid email without domain",
			input:         "test@.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			name:          "invalid email without @ symbol",
			input:         "testexample.com",
			expectedError: errors.New("userEmail format is invalid"),
			expectedValue: nil,
		},
		{
			name:          "overlength email",
			input:         "a" + string(make([]rune, 255)) + "@example.com",
			expectedError: errors.New("userEmail length over 256"),
			expectedValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewUserEmail(tt.input)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("got error %v, want %v", err, tt.expectedError)
			}
			if tt.expectedValue != nil && result.String() != tt.expectedValue.String() {
				t.Errorf("got result %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
