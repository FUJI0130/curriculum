package mentorapp_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/mentorrecruitmentdm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	mockCategory "github.com/FUJI0130/curriculum/src/core/mock/mock_category"
	mockMentorRecruitmentTag "github.com/FUJI0130/curriculum/src/core/mock/mock_mentor_recruitment_tag"
	mockMentorRecruitment "github.com/FUJI0130/curriculum/src/core/mock/mock_mentorrecruitment"
	mockTag "github.com/FUJI0130/curriculum/src/core/mock/mock_tag"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Define more test cases as needed
func TestCreateMentorRecruitmentAppService_Exec(t *testing.T) {
	ctx := context.TODO()
	mockCategoryName := "Test Category"

	tests := []struct {
		title    string
		input    *mentorapp.CreateMentorRecruitmentRequest
		mockFunc func(mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository, mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository, mockTagRepo *mockTag.MockTagRepository, mockCategoryRepo *mockCategory.MockCategoryRepository)
		wantErr  error
	}{

		{
			title: "カテゴリが存在しない場合のメンター募集作成と保存",
			input: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "テストメンター募集",
				CategoryName:          mockCategoryName,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0), // 1ヶ月後
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "メンター募集の詳細説明",
				Status:                1,
				TagNames:              []string{"Tag1", "Tag2"},
			},
			mockFunc: func(mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository, mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository, mockTagRepo *mockTag.MockTagRepository, mockCategoryRepo *mockCategory.MockCategoryRepository) {
				// カテゴリが存在しないことをモック（sql.ErrNoRows を返すように変更）
				mockCategoryRepo.EXPECT().FindByName(gomock.Any(), mockCategoryName).Return(nil, sql.ErrNoRows)

				// 新しいカテゴリの保存をモック
				mockCategoryRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)

				// タグの処理に関するモック設定
				mockTagRepo.EXPECT().FindByName(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
				mockTagRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)

				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, mentorRecruitmentTag *mentorrecruitmentdm.MentorRecruitmentTag) error {
						assert.NotNil(t, mentorRecruitmentTag)
						return nil
					}).AnyTimes()
			},
			wantErr: nil,
		},
		{
			title: "タグが存在しない場合の新しいタグの生成と保存",
			input: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "新規メンター募集",
				CategoryName:          mockCategoryName,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0), // 1ヶ月後
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "メンター募集の詳細説明",
				Status:                1,
				TagNames:              []string{"New Tag"},
			},
			mockFunc: func(mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository, mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository, mockTagRepo *mockTag.MockTagRepository, mockCategoryRepo *mockCategory.MockCategoryRepository) {
				// カテゴリが存在することをモック
				mockCategoryRepo.EXPECT().FindByName(ctx, mockCategoryName).Return(&categorydm.Category{}, nil)

				// タグが存在しないことをモック
				mockTagRepo.EXPECT().FindByName(ctx, "New Tag").Return(nil, nil)
				// 新しいタグの保存をモック
				mockTagRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, mentorRecruitmentTag *mentorrecruitmentdm.MentorRecruitmentTag) error {
						assert.NotNil(t, mentorRecruitmentTag)
						return nil
					}).AnyTimes()
			},
			wantErr: nil,
		},
		{
			title: "既存のタグを使用するケース",
			input: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "既存タグ使用のメンター募集",
				CategoryName:          mockCategoryName,
				BudgetFrom:            20000,
				BudgetTo:              40000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0), // 1ヶ月後
				ConsultationFormat:    2,
				ConsultationMethod:    2,
				Description:           "既存タグを使用するメンター募集の説明",
				Status:                1,
				TagNames:              []string{"Existing Tag"},
			},
			mockFunc: func(mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository, mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository, mockTagRepo *mockTag.MockTagRepository, mockCategoryRepo *mockCategory.MockCategoryRepository) {
				// カテゴリと既存のタグが存在することをモック
				mockCategoryRepo.EXPECT().FindByName(ctx, mockCategoryName).Return(&categorydm.Category{}, nil)
				existingTag, _ := tagdm.GenWhenCreateTag("Existing Tag")
				mockTagRepo.EXPECT().FindByName(ctx, "Existing Tag").Return(existingTag, nil)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, mentorRecruitmentTag *mentorrecruitmentdm.MentorRecruitmentTag) error {
						assert.NotNil(t, mentorRecruitmentTag)
						return nil
					}).AnyTimes()
			},
			wantErr: nil,
		},
		{
			title: "メンター募集の正常な作成と保存",
			input: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "新しいメンター募集",
				CategoryName:          mockCategoryName,
				BudgetFrom:            30000,
				BudgetTo:              60000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 2, 0), // 2ヶ月後
				ConsultationFormat:    3,
				ConsultationMethod:    3,
				Description:           "新しいメンター募集の詳細説明",
				Status:                1,
				TagNames:              []string{"Existing Tag"},
			},
			mockFunc: func(mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository, mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository, mockTagRepo *mockTag.MockTagRepository, mockCategoryRepo *mockCategory.MockCategoryRepository) {
				// カテゴリと既存のタグが存在することをモック
				mockCategoryRepo.EXPECT().FindByName(ctx, mockCategoryName).Return(&categorydm.Category{}, nil)
				existingTag, _ := tagdm.GenWhenCreateTag("Existing Tag")
				mockTagRepo.EXPECT().FindByName(ctx, "Existing Tag").Return(existingTag, nil)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, mentorRecruitmentTag *mentorrecruitmentdm.MentorRecruitmentTag) error {
						assert.NotNil(t, mentorRecruitmentTag)
						return nil
					}).AnyTimes()
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMentorRecruitmentRepo := mockMentorRecruitment.NewMockMentorRecruitmentRepository(ctrl)
			mockMentorRecruitmentTagRepo := mockMentorRecruitmentTag.NewMockMentorRecruitmentsTagsRepository(ctrl)
			mockTagRepo := mockTag.NewMockTagRepository(ctrl)
			mockCategoryRepo := mockCategory.NewMockCategoryRepository(ctrl)

			app := mentorapp.NewCreateMentorRecruitmentAppService(mockMentorRecruitmentRepo, mockMentorRecruitmentTagRepo, mockTagRepo, mockCategoryRepo)
			tt.mockFunc(mockMentorRecruitmentRepo, mockMentorRecruitmentTagRepo, mockTagRepo, mockCategoryRepo)

			err := app.Exec(ctx, tt.input)

			// Assertions here
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
