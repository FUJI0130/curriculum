package userdm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	// createdAtとupdatedAtの作成

	//●error処理を書く
	createdAt, err := NewUserCreatedAt(time.Now())
	if err != nil {
		t.Fatalf("failed to create createdAt: %v", err) // テストを失敗させてエラーメッセージを表示
	}

	updatedAt, err := NewUserUpdatedAt(time.Now())
	if err != nil {
		t.Fatalf("failed to create updatedAt: %v", err) // テストを失敗させてエラーメッセージを表示
	}

	tests := []struct {
		name        string
		email       string
		password    string
		profile     string
		createdAt   *UserCreatedAt
		updatedAt   *UserUpdatedAt
		expectError bool // expectErrorを追加
	}{
		{
			name:        "TestName",
			email:       "test@example.com",
			password:    "Password",
			profile:     "Profile",
			createdAt:   createdAt,
			updatedAt:   updatedAt,
			expectError: false, // テストケースに合わせてここを設定
		},
		// ここに他のテストケースを追加出来る
	}

	for _, test := range tests {
		user, err := NewUser(test.name, test.email, test.password, test.profile, test.createdAt.DateTime(), test.updatedAt.DateTime())

		if test.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, user.name, user.name)
			assert.Equal(t, user.email, user.email)
			assert.Equal(t, user.password, user.password)
			assert.Equal(t, user.profile, user.profile)
			assert.Equal(t, user.createdAt, user.createdAt)
			assert.Equal(t, user.updatedAt, user.updatedAt)
		}
	}
}
