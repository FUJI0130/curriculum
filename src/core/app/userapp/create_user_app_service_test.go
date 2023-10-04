package userapp_test

import (
	"context"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	domainErrors "github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	mockExistByNameDomainService "github.com/FUJI0130/curriculum/src/core/mock/mockExistByNameDomainService"
	"github.com/FUJI0130/curriculum/src/core/mock/mockTag"
	"github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	mockName := "Test User"
	mockEmail := "test@example.com"
	mockPassword := "newuserpassword"
	mockProfile := "Sample Profile"
	mockTagName := "Test Tag"

	ctx := context.TODO()

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
				Skills:   []userapp.SkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
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
			wantErr: domainErrors.ErrUserNameAlreadyExists(nil, mockName, "TestCreateUserAppService_Exec 存在するユーザーを作成"),
		},
		{
			title: "タグの新規作成",
			input: &userapp.CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []userapp.SkillRequest{{TagName: "New Tag", Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
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
				Skills:   []userapp.SkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []userapp.CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
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
				Skills:   []userapp.SkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}, {TagName: mockTagName, Evaluation: 4, Years: 1}},
				Profile:  mockProfile,
				Careers:  []userapp.CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockExistService.EXPECT().Exec(ctx, mockName).Return(false, nil)

				existingTag, _ := tagdm.GenWhenCreateTag(mockTagName)
				mockTagRepo.EXPECT().FindByNames(ctx, []string{mockTagName, mockTagName}).Return([]*tagdm.Tag{existingTag, existingTag}, nil).Times(1)

			},
			wantErr: customerrors.ErrDuplicateSkillTag(nil, mockTagName, "TestCreateUserAppService_Exec  ユーザーが同じスキルタグを複数持つ場合"), // 期待されるエラーメッセージ
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockUserRepo := mockUser.NewMockUserRepository(ctrl)
			mockTagRepo := mockTag.NewMockTagRepository(ctrl)
			mockExistService := mockExistByNameDomainService.NewMockExistByNameDomainService(ctrl)
			app := userapp.NewCreateUserAppService(mockUserRepo, mockTagRepo, mockExistService)
			tt.mockFunc(mockUserRepo, mockTagRepo, mockExistService)

			err := app.Exec(context.TODO(), tt.input)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
