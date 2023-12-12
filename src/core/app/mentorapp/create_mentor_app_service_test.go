package mentorapp_test

import (
	"context"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/mentorapp"
	"github.com/FUJI0130/curriculum/src/core/domain/categorydm"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	mockCategory "github.com/FUJI0130/curriculum/src/core/mock/mock_category"
	mockMentorRecruitmentTag "github.com/FUJI0130/curriculum/src/core/mock/mock_mentor_recruitment_tag"
	mockMentorRecruitment "github.com/FUJI0130/curriculum/src/core/mock/mock_mentorrecruitment"
	mockTagdm "github.com/FUJI0130/curriculum/src/core/mock/mock_tagdm"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateMentorRecruitmentAppService_Exec_CategoryNotFound(t *testing.T) {
	mockCategoryID := "12345678-1234-56ab-7c89-1011d2e34fgh"
	mockNewTagName := "NewTag"
	mockExistingTagID := "e5431b9c-6212-4874-ac10-cc6209c96246"
	mockTagID := "12345678-1234-56ab-7c89-1011d2e34fgh"

	tests := []struct {
		name      string
		request   *mentorapp.CreateMentorRecruitmentRequest
		mockSetup func(
			mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository,
			mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository,
			mockCategoryRepo *mockCategory.MockCategoryRepository,
			mockTagDomainService *mockTagdm.MockTagDomainService,
		)
		wantErr error
	}{
		{
			name: "CategoryNotFound",
			request: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "Test Mentor Recruitment",
				CategoryID:            mockCategoryID,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0),
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "Detailed description of the mentor recruitment",
				Status:                1,
				TagNames:              []string{"Tag1", "Tag2"},
			},
			mockSetup: func(
				mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository,
				mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository,
				mockCategoryRepo *mockCategory.MockCategoryRepository,
				mockTagDomainService *mockTagdm.MockTagDomainService,
			) {
				// Setup mock for category not found
				mockCategoryRepo.EXPECT().
					FindByID(gomock.Any(), mockCategoryID).
					Return(nil, customerrors.NewNotFoundError("カテゴリが見つかりません")).
					Times(1)
			},
			wantErr: customerrors.NewNotFoundError("カテゴリが見つかりません"),
		},
		{
			name: "NewTagCreationAndSave",
			request: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "New Mentor Recruitment",
				CategoryID:            mockCategoryID,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0),
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "Detailed description of the mentor recruitment",
				Status:                1,
				TagIDs:                []string{""}, // 空のTagIDを設定
				TagNames:              []string{mockNewTagName},
			},
			mockSetup: func(
				mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository,
				mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository,
				mockCategoryRepo *mockCategory.MockCategoryRepository,
				mockTagDomainService *mockTagdm.MockTagDomainService,
			) {
				// カテゴリが存在することをモック
				mockCategoryRepo.EXPECT().
					FindByID(gomock.Any(), mockCategoryID).
					Return(&categorydm.Category{}, nil).
					Times(1)

				// 新規タグの生成をモック
				mockTagDomainService.EXPECT().
					ProcessTag(gomock.Any(), gomock.Eq(""), gomock.Eq(mockNewTagName)).
					Return(&tagdm.Tag{}, nil).
					Times(1)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)

				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			wantErr: nil,
		},
		{
			name: "ExistingTagUse",
			request: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "New Mentor Recruitment",
				CategoryID:            mockCategoryID,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0),
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "Detailed description of the mentor recruitment",
				Status:                1,
				TagIDs:                []string{mockExistingTagID}, // 既存のタグIDを設定
				TagNames:              []string{},
			},
			mockSetup: func(
				mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository,
				mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository,
				mockCategoryRepo *mockCategory.MockCategoryRepository,
				mockTagDomainService *mockTagdm.MockTagDomainService,
			) {
				// カテゴリが存在することをモック
				mockCategoryRepo.EXPECT().
					FindByID(gomock.Any(), mockCategoryID).
					Return(&categorydm.Category{}, nil).
					Times(1)

				// 既存のタグが取得されることをモック
				mockTagDomainService.EXPECT().
					ProcessTag(gomock.Any(), gomock.Eq(mockExistingTagID), gomock.Any()).
					Return(&tagdm.Tag{}, nil).
					Times(1)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)

				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			wantErr: nil,
		},
		{
			name: "SuccessfulCreationAndSave",
			request: &mentorapp.CreateMentorRecruitmentRequest{
				Title:                 "Successful Mentor Recruitment",
				CategoryID:            mockCategoryID,
				BudgetFrom:            10000,
				BudgetTo:              50000,
				ApplicationPeriodFrom: time.Now(),
				ApplicationPeriodTo:   time.Now().AddDate(0, 1, 0),
				ConsultationFormat:    1,
				ConsultationMethod:    1,
				Description:           "Detailed description of the mentor recruitment",
				Status:                1,
				TagIDs:                []string{mockTagID}, // 既存のタグID
				TagNames:              []string{},
			},
			mockSetup: func(
				mockMentorRecruitmentRepo *mockMentorRecruitment.MockMentorRecruitmentRepository,
				mockMentorRecruitmentTagRepo *mockMentorRecruitmentTag.MockMentorRecruitmentsTagsRepository,
				mockCategoryRepo *mockCategory.MockCategoryRepository,
				mockTagDomainService *mockTagdm.MockTagDomainService,
			) {
				// カテゴリが存在することをモック
				mockCategoryRepo.EXPECT().
					FindByID(gomock.Any(), mockCategoryID).
					Return(&categorydm.Category{}, nil).
					Times(1)

				// 既存のタグが取得されることをモック
				mockTagDomainService.EXPECT().
					ProcessTag(gomock.Any(), gomock.Eq(mockTagID), gomock.Any()).
					Return(&tagdm.Tag{}, nil).
					Times(1)

				// メンター募集の保存をモック
				mockMentorRecruitmentRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)

				// 中間テーブルの更新をモック
				mockMentorRecruitmentTagRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMentorRecruitmentRepo := mockMentorRecruitment.NewMockMentorRecruitmentRepository(ctrl)
			mockMentorRecruitmentTagRepo := mockMentorRecruitmentTag.NewMockMentorRecruitmentsTagsRepository(ctrl)
			mockCategoryRepo := mockCategory.NewMockCategoryRepository(ctrl)
			mockTagDomainService := mockTagdm.NewMockTagDomainService(ctrl)

			app := mentorapp.NewCreateMentorRecruitmentAppService(
				mockMentorRecruitmentRepo,
				mockMentorRecruitmentTagRepo,
				mockCategoryRepo,
				mockTagDomainService,
			)

			tt.mockSetup(mockMentorRecruitmentRepo, mockMentorRecruitmentTagRepo, mockCategoryRepo, mockTagDomainService)

			err := app.Exec(context.TODO(), tt.request)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
