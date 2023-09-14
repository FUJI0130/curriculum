package userapp

import (
	"context"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/mock/mockTag"
	"github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	// Set up common data for all tests
	mockName := "Test User"
	mockEmail := "test@example.com"
	mockPassword := "newuserpassword"
	mockProfile := "Sample Profile"
	mockTagName := "Test Tag"

	ctx := context.TODO()

	tests := []struct {
		name     string
		input    *CreateUserRequest
		mockFunc func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockUser.MockExistByNameDomainService)
		wantErr  error
	}{
		{
			name: "ユーザー新規作成",
			input: &CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []SkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockUser.MockExistByNameDomainService) {

				mockExistService.EXPECT().IsExist(ctx, mockName).Return(false, nil).Times(1)
				mockTagRepo.EXPECT().FindByName(ctx, mockTagName).Return(nil, tagdm.ErrTagNotFound).Times(1)
				mockTagRepo.EXPECT().CreateNewTag(ctx, mockTagName).Return(&tagdm.Tag{}, nil).Times(1)
				mockUserRepo.EXPECT().Store(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

			},
			wantErr: nil,
		},
		{
			name: "存在するユーザーを作成",
			input: &CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockUser.MockExistByNameDomainService) {
				mockExistService.EXPECT().IsExist(ctx, mockName).Return(true, nil)
			},
			wantErr: ErrUserNameAlreadyExists,
		},
		// Add more test cases as needed
		{
			name: "タグの新規作成",
			input: &CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []SkillRequest{{TagName: "New Tag", Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockUser.MockExistByNameDomainService) {
				mockExistService.EXPECT().IsExist(ctx, mockName).Return(false, nil)
				mockTagRepo.EXPECT().FindByName(ctx, "New Tag").Return(nil, tagdm.ErrTagNotFound)
				mockTagRepo.EXPECT().CreateNewTag(ctx, "New Tag").Return(&tagdm.Tag{}, nil)
				mockUserRepo.EXPECT().Store(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "既存のタグを使用してユーザーを作成",
			input: &CreateUserRequest{
				Name:     mockName,
				Email:    mockEmail,
				Password: mockPassword,
				Skills:   []SkillRequest{{TagName: mockTagName, Evaluation: 5, Years: 2}},
				Profile:  mockProfile,
				Careers:  []CareersRequest{{Detail: "Dev", AdFrom: time.Now(), AdTo: time.Now().AddDate(1, 0, 0)}},
			},
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository, mockExistService *mockUser.MockExistByNameDomainService) {
				mockExistService.EXPECT().IsExist(ctx, mockName).Return(false, nil)

				// 既存のタグオブジェクトを作成
				existingTag, _ := tagdm.NewTag(mockTagName)
				mockTagRepo.EXPECT().FindByName(ctx, mockTagName).Return(existingTag, nil)
				mockUserRepo.EXPECT().Store(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUserRepo := mockUser.NewMockUserRepository(ctrl)
			mockTagRepo := mockTag.NewMockTagRepository(ctrl)
			mockExistService := mockUser.NewMockExistByNameDomainService(ctrl)
			app := NewCreateUserAppService(mockUserRepo, mockTagRepo, mockExistService)
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
