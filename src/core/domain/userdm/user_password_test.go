package userdm

import (
	"testing"
)

func TestNewUserPassword(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "empty password",
			input:         "",
			expectedError: "userPassword cannot be empty",
		},
		{
			name:          "valid password",
			input:         "AValidPassword123",
			expectedError: "",
		},
		{
			name:          "too short password",
			input:         "Short1",
			expectedError: "userPassword length under 12",
		},
		{
			name:          "too long password",
			input:         "ThisIsAReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyLongPasswordThatIsDefinitelyOver256CharactersLongAndShouldReturnAnErrorBecauseOfItsLength1234567890LetsMakeThisStringExactlyTwoHundredFiftyFiveCharactersLongAddingMoreNow+",
			expectedError: "userPassword length over 256",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUserPassword(tt.input)
			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedError)
				} else if err.Error() != tt.expectedError {
					t.Errorf("expected error %v, got %v", tt.expectedError, err.Error())
				}
			} else if err != nil {
				t.Errorf("expected no error, got %v", err.Error())
			}
		})
	}
}
