package userdm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUserID_check(t *testing.T) {
	// testUserID2とtestUserID2を作成
	testUserID1, err := NewUserID()
	if err != nil {
		t.Fatalf("Failed to create userID1: %v", err)
	}

	testUserID2, err := NewUserID()
	if err != nil {
		t.Fatalf("Failed to create userID2: %v", err)
	}

	// テストケースを定義
	tests := []struct {
		name  string
		id1   *UserID
		id2   *UserID
		equal bool
	}{
		{
			name:  "testUserID1 and testUserID2 should not be equal",
			id1:   testUserID1,
			id2:   testUserID2,
			equal: false,
		},
		{
			name:  "testUserID1 and itself should be equal",
			id1:   testUserID1,
			id2:   testUserID1,
			equal: true,
		},
		{
			name:  "testUserID2 and itself should be equal",
			id1:   testUserID2,
			id2:   testUserID2,
			equal: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// t.Logf("Comparing UUIDs: %v and %v", tt.id1.String(), tt.id2.String())

			equal := tt.id1.Equal(tt.id2)
			if tt.equal {
				assert.True(t, equal, "%v should be equal to %v", tt.id1, tt.id2)
			} else {
				assert.False(t, equal, "%v should not be equal to %v", tt.id1, tt.id2)
			}

			// オプション：ユーザIDを出力
			t.Logf("UserID1: %v", tt.id1.String())
			t.Logf("UserID2: %v", tt.id2.String())
		})
	}
}

func TestNewUser(t *testing.T) {
	// createdAtとupdatedAtの作成
	createdAt, _ := NewUserCreatedAt(time.Now())
	updatedAt, _ := NewUserUpdatedAt(time.Now())

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
			expectError: false, // テストケースに合わせてここを設定します
		},
		// ここに他のテストケースを追加できます。
	}

	for _, test := range tests {
		user, err := NewUser(test.name, test.email, test.password, test.profile, test.createdAt.DateTime(), test.updatedAt.DateTime())

		if test.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			// ここでは値オブジェクトの比較が必要かもしれませんが、
			// その具体的な方法は値オブジェクトの実装によります
			assert.Equal(t, user.name, user.name)
			assert.Equal(t, user.email, user.email)
			assert.Equal(t, user.password, user.password)
			assert.Equal(t, user.profile, user.profile)
			assert.Equal(t, user.createdAt, user.createdAt)
			assert.Equal(t, user.updatedAt, user.updatedAt)
		}
	}
}
