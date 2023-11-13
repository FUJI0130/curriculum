package userapp_test

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	mockExistByNameDomainService "github.com/FUJI0130/curriculum/src/core/mock/mock_exist_by_name_domain_service"
	mockTag "github.com/FUJI0130/curriculum/src/core/mock/mock_tag"
	mockUser "github.com/FUJI0130/curriculum/src/core/mock/mock_user"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	mockName := "Test User"
	mockEmail := "test@example.com"
	mockPassword := "newuserpassword"
	mockProfile := "Sample Profile"
	mockTagName := "Test Tag"

	// ctx := context.TODO()
	ctx := context.WithValue(context.Background(), "Conn", &sqlx.Tx{})

	tests := []struct {
		title    string
		input    *userapp.CreateUserRequest
		mockFunc func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService)
		wantErr  error
	}{
		{
			title: "ユーザー新規作成",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []userapp.CreateSkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CreateCareerRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {

				mockExistService.EXPECT().Exec(ctx, mockName).Return(false, nil).Times(1)
				mockTagRepo.EXPECT().FindByNames(ctx, []string{mockTagName}).Return([]*tagdm.Tag{}, nil).Times(1)
				mockTagRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil).Times(1)
				mockUserRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			title: "存在するユーザーを作成",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockExistService.EXPECT().Exec(ctx, mockName).Return(true, nil)
			},
			wantErr: customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec UserName isExist  name is : %s", mockName),
		},
		{
			title: "タグの新規作成",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []userapp.CreateSkillRequest{{TagName: "New Tag", Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CreateCareerRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockExistService.EXPECT().Exec(ctx, mockName).Return(false, nil)
				mockTagRepo.EXPECT().FindByNames(ctx, []string{"New Tag"}).Return([]*tagdm.Tag{}, nil).Times(1)

				mockTagRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil).Times(1)
				mockUserRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			title: "既存のタグを使用してユーザーを作成",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []userapp.CreateSkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CreateCareerRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockExistService.EXPECT().Exec(ctx, mockName).Return(false, nil)

				existingTag, _ := tagdm.GenWhenCreateTag(mockTagName)
				mockTagRepo.EXPECT().FindByNames(ctx, []string{mockTagName}).Return([]*tagdm.Tag{existingTag}, nil).Times(1)

				mockUserRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			title: "ユーザーが同じスキルタグを複数回持つ場合",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []userapp.CreateSkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}, {TagName: mockTagName, Evaluation: 4, Years: 1}},
				Profile:  mockProfile,
				Careers:  []userapp.CreateCareerRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockExistService.EXPECT().Exec(ctx, mockName).Return(false, nil)
				expectedTagNames := []string{mockTagName}
				log.Printf("Expected tagNames in mock setting: %v", expectedTagNames)
				// 実際のメソッド呼び出しに合わせたモック設定
				// mockTagRepo.EXPECT().FindByNames(ctx, gomock.Eq([]string{mockTagName})).Return([]*tagdm.Tag{}, nil).Times(1)
				// mockTagRepo.EXPECT().FindByNames(ctx, []string{mockTagName}).Return([]*tagdm.Tag{}, nil).Times(1)
				mockTagRepo.EXPECT().FindByNames(gomock.Any(), gomock.Eq([]string{mockTagName})).Return([]*tagdm.Tag{}, nil).Times(1)

			},
			wantErr: customerrors.NewUnprocessableEntityErrorf("Create_user_app_service  Exec tagname is : %s", mockTagName),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			// t.Parallel()

			ctrl := gomock.NewController(t)
			mockUserRepo := mockUser.NewMockUserRepository(ctrl)
			mockTagRepo := mockTag.NewMockTagRepository(ctrl)
			mockExistService := mockExistByNameDomainService.NewMockExistByNameDomainService(ctrl)
			app := userapp.NewCreateUserAppService(mockUserRepo, mockTagRepo, mockExistService)
			tt.mockFunc(mockUserRepo, mockTagRepo, mockExistService)

			log.Printf("TestCode: ctx is : %v", ctx)
			//ctxの型を確認する
			log.Printf("TestCode: ctx type is : %T", ctx)
			if conn, ok := ctx.Value("Conn").(*sqlx.Tx); ok {
				log.Printf("TestCode: ctx contains a transaction object: %#v", conn)
			} else {
				log.Printf("TestCode: ctx does not contain the expected transaction object")
			}

			err := app.Exec(ctx, tt.input)
			if tt.wantErr != nil {
				var unprocessableEntityErr *customerrors.UnprocessableEntityErrorType
				if errors.As(err, &unprocessableEntityErr) {
					if strings.Contains(unprocessableEntityErr.Error(), "UserName isExist") {
						// ユーザー名が存在する場合のエラーメッセージの検証
						assert.Contains(t, unprocessableEntityErr.Error(), "Create_user_app_service  Exec UserName isExist")
					} else if strings.Contains(unprocessableEntityErr.Error(), "Skill with tag name") {
						// スキルタグの重複に関するエラーメッセージの検証
						assert.Contains(t, unprocessableEntityErr.Error(), "Skill with tag name")
					} else {
						t.Errorf("Unexpected unprocessable entity error message: %v", unprocessableEntityErr.Error())
					}
				} else {
					t.Errorf("Expected an unprocessable entity error, got %v", err)
				}
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
