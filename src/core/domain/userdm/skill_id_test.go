package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkillID_check(t *testing.T) {
	testSkillID1, err := NewSkillID()
	if err != nil {
		t.Fatalf("Failed to create testSkillID1: %v", err)
	}

	testSkillID2, err := NewSkillID()
	if err != nil {
		t.Fatalf("Failed to create testSkillID2: %v", err)
	}

	// テストケースを定義
	tests := []struct {
		name  string
		id1   SkillID
		id2   SkillID
		isErr bool
	}{
		{
			name:  "testSkillID1 と testSkillID2 は等しくない事を確認するテスト",
			id1:   testSkillID1,
			id2:   testSkillID2,
			isErr: false,
		},
		{
			name:  "testSkillID1 自身を比べて等しい事を確認するテスト",
			id1:   testSkillID1,
			id2:   testSkillID1,
			isErr: true,
		},
		{
			name:  "testSkillID2 自身を比べて等しい事を確認するテスト",
			id1:   testSkillID2,
			id2:   testSkillID2,
			isErr: true,
		},
	}

	// tt = 現在のループでのtestcase
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// tt.id1 tt.id2が等しいかどうかをチェックして、結果を格納
			equal := tt.id1.Equal(tt.id2)
			if tt.isErr {
				assert.True(t, equal, "%v should be equal to %v", tt.id1, tt.id2)
			} else {
				assert.False(t, equal, "%v should not be equal to %v", tt.id1, tt.id2)
			}
		})
	}
}
