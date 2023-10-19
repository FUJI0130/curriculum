package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserPassword(t *testing.T) {
	tests := []struct {
		title         string
		input         string
		expectedError string
	}{
		{
			title:         "パスワードが空",
			input:         "",
			expectedError: "userPassword cannot be empty",
		},
		{
			title:         "有効なパスワード",
			input:         "AValidPassword123",
			expectedError: "",
		},
		{
			title:         "パスワードが12文字以下",
			input:         "Short1234567",
			expectedError: "userPassword length under 12",
		},
		{
			title:         "パスワードが256文字以上",
			input:         "ThisIsAReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyReallyLongPasswordThatIsDefinitelyOver256CharactersLongAndShouldReturnAnErrorBecauseOfItsLength1234567890LetsMakeThisStringExactlyTwoHundredFiftyFiveCharactersLongAddingMoreNow++",
			expectedError: "userPassword length over nameMaxlength",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			_, err := NewUserPassword(tt.input)
			if tt.expectedError != "" {
				assert.Error(t, err, "エラーが期待されましたが、nilが返されました")
				assert.EqualError(t, err, tt.expectedError, "エラーメッセージが一致しません")
				return
			}

			assert.NoError(t, err, "エラーは期待されませんでしたが、%vが返されました", err)
		})
	}

}
