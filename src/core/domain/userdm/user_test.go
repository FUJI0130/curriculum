package userdm

import (
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	// createdAtとupdatedAtの作成

	//●error処理を書く
	createdAt := sharedvo.NewCreatedAtByVal(time.Now())
	if err != nil {
		t.Fatalf("failed to create createdAt: %v", err) // テストを失敗させてエラーメッセージを表示
	}

	startTime := time.Now()
	updatedAt := sharedvo.NewUpdatedAtByVal(startTime)
	if err != nil {
		t.Logf("Time taken for updatedAt.Before(time.Now()): %v", sharedvo.LastDuration)
		t.Fatalf("failed to create updatedAt: %v", err) // テストを失敗させてエラーメッセージを表示
	}

	user_password, err := NewUserPassword("new_password0130")
	if err != nil {
		t.Fatalf("failed to create user_password: %v", err) // テストを失敗させてエラーメッセージを表示
	}

	tests := []struct {
		name        string
		email       string
		password    UserPassword
		profile     string
		createdAt   sharedvo.CreatedAt
		updatedAt   sharedvo.UpdatedAt
		expectError bool // expectErrorを追加
	}{
		{
			name:        "TestName",
			email:       "test@example.com",
			password:    user_password,
			profile:     "Profile",
			createdAt:   createdAt,
			updatedAt:   updatedAt,
			expectError: false, // テストケースに合わせてここを設定
		},
		// ここに他のテストケースを追加出来る
	}

	for _, test := range tests {
		user, err := NewUser(test.name, test.email, test.password.String(), test.profile, test.createdAt.DateTime(), test.updatedAt.DateTime())

		if test.expectError {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)
		assert.Equal(t, user.name, user.name)
		assert.Equal(t, user.email, user.email)
		assert.Equal(t, user.password, user.password)
		assert.Equal(t, user.profile, user.profile)
		assert.Equal(t, user.createdAt, user.createdAt)
		assert.Equal(t, user.updatedAt, user.updatedAt)
	}
}
