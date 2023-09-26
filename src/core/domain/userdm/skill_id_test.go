package userdm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkillID_check(t *testing.T) {

	// テストケースを定義
	tests := []struct {
		title   string
		idGen   func() (SkillID, error) // SkillIDとerrorを返す関数を追加
		wantErr bool                    // エラーが発生することを期待するかのフラグ
		isErr   bool                    // id1とid2が等しいかどうか
	}{
		{
			title:   "testSkillID1 と testSkillID2 は等しくない事を確認するテスト",
			idGen:   NewSkillID,
			wantErr: false,
			isErr:   false,
		},
		{
			title:   "testSkillID1 自身を比べて等しい事を確認するテスト",
			idGen:   NewSkillID,
			wantErr: false,
			isErr:   true,
		},
		{
			title:   "testSkillID2 自身を比べて等しい事を確認するテスト",
			idGen:   NewSkillID,
			wantErr: false,
			isErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			id1, err1 := tt.idGen()

			// 自身と比べるテストケースの場合、id1をid2にコピー
			var id2 SkillID
			if tt.isErr {
				id2 = id1
			} else {
				var err2 error
				id2, err2 = tt.idGen()
				assert.NoError(t, err2)
			}

			assert.NoError(t, err1)

			// id1とid2が等しいかどうかをチェックして、結果を格納
			equal := id1.Equal(id2)
			if tt.isErr {
				assert.True(t, equal, "%v should be equal to %v", id1, id2)
			} else {
				assert.False(t, equal, "%v should not be equal to %v", id1, id2)
			}
		})
	}
}
