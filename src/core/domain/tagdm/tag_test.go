// src/core/domain/tagdm/tag_test.go

package tagdm_test

import (
	"context"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/mock/mockTag" // ここでモックをインポート
	"github.com/golang/mock/gomock"
)

func TestTagRepositoryMethods(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTagRepo := mockTag.NewMockTagRepository(ctrl)
	ctx := context.TODO()

	tests := []struct {
		name          string
		tagName       string
		storeErr      error
		findByNameErr error
		findByIDErr   error
	}{
		{
			name:          "Successful flow",
			tagName:       "testTag",
			storeErr:      nil,
			findByNameErr: nil,
			findByIDErr:   nil,
		},
		// 追加テストケース
		{
			name:          "Failed due to empty tag name",
			tagName:       "",
			storeErr:      tagdm.ErrTagNameEmpty,
			findByNameErr: tagdm.ErrTagNameEmpty,
			findByIDErr:   tagdm.ErrTagNameEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := tagdm.NewTag(tt.tagName, time.Now(), time.Now())
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Store mock
			mockTagRepo.EXPECT().Store(ctx, tag).Return(tt.storeErr)

			// FindByName mock
			mockTagRepo.EXPECT().FindByName(ctx, tt.tagName).Return(tag, tt.findByNameErr)

			// FindByID mock
			mockTagRepo.EXPECT().FindByID(ctx, string(tag.ID())).Return(tag, tt.findByIDErr)

			// ここで実際のテストロジックを書く。上記はモックの期待値を設定するだけです。

			if err := mockTagRepo.Store(ctx, tag); err != tt.storeErr {
				t.Fatalf("Expected storeErr: %v, got: %v", tt.storeErr, err)
			}

			foundTag, err := mockTagRepo.FindByName(ctx, tt.tagName)
			if err != tt.findByNameErr || foundTag.Name() != tt.tagName {
				t.Fatalf("Expected findByNameErr: %v, got: %v", tt.findByNameErr, err)
			}

			foundTagByID, err := mockTagRepo.FindByID(ctx, string(tag.ID()))
			if err != tt.findByIDErr || foundTagByID.ID() != tag.ID() {
				t.Fatalf("Expected findByIDErr: %v, got: %v", tt.findByIDErr, err)
			}
		})
	}
}
