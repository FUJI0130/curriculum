package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				if assert.NotNil(t, err, "エラーが期待されましたが、nilが返されました") {
					assert.Equal(t, tt.expectedError, err.Error(), "エラーメッセージが一致しません")
				}
				return
			}
			assert.Nil(t, err, "エラーは期待されませんでしたが、%vが返されました", err)
		})
	}

}
