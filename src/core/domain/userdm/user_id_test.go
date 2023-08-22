package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserID_check(t *testing.T) {
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
		id1   UserID
		id2   UserID
		isErr bool
	}{
		{
			name:  "testUserID1 と testUserID2 は等しくない事を確認するテスト",
			id1:   testUserID1,
			id2:   testUserID2,
			isErr: false,
		},
		{
			name:  "testUserID1 自身を比べて等しい事を確認するテスト",
			id1:   testUserID1,
			id2:   testUserID1,
			isErr: true,
		},
		{
			name:  "testUserID2  自身を比べて等しい事を確認するテスト",
			id1:   testUserID2,
			id2:   testUserID2,
			isErr: true,
		},
	}

	//tt= 現在のループでのtestcase
	for _, tt := range tests {
		//テストケースのサブテストを実行するための関数 第１引数：サブテストの名前　第２引数：実際のテストのコード（関数)
		t.Run(tt.name, func(t *testing.T) {

			//tt.id1 tt.id2が等しいかどうかをチェックして、結果を格納
			equal := tt.id1.Equal(tt.id2)
			if tt.isErr {
				assert.True(t, equal, "%v should be equal to %v", tt.id1, tt.id2)
			} else {
				assert.False(t, equal, "%v should not be equal to %v", tt.id1, tt.id2)
			}
		})
	}
}
