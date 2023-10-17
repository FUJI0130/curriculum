package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserID_check(t *testing.T) {
	tests := []struct {
		title string
		id1   UserID
		id2   UserID
		isErr bool
	}{
		{
			title: "testUserID1 と testUserID2 は等しくない事を確認するテスト",
			isErr: false,
		},
		{
			title: "testUserID1 自身を比べて等しい事を確認するテスト",
			isErr: true,
		},
		{
			title: "testUserID2 自身を比べて等しい事を確認するテスト",
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			testUserID1, err := NewUserID()
			if err != nil {
				t.Fatalf("Failed to create userID1: %v", err)
			}

			testUserID2, err := NewUserID()
			if err != nil {
				t.Fatalf("Failed to create userID2: %v", err)
			}

			var compare1, compare2 UserID
			switch tt.title {
			case "testUserID1 と testUserID2 は等しくない事を確認するテスト":
				compare1 = testUserID1
				compare2 = testUserID2
			case "testUserID1 自身を比べて等しい事を確認するテスト":
				compare1 = testUserID1
				compare2 = testUserID1
			case "testUserID2 自身を比べて等しい事を確認するテスト":
				compare1 = testUserID2
				compare2 = testUserID2
			}

			equal := compare1.Equal(compare2)
			if tt.isErr {
				assert.True(t, equal, "%v should be equal to %v", compare1, compare2)
			} else {
				assert.False(t, equal, "%v should not be equal to %v", compare1, compare2)
			}
		})
	}
}
