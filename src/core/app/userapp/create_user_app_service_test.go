package userapp

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/mock/mockTag"
	"github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	type testCase struct {
		name     string
		request  *CreateUserRequest
		wantErr  error
		mockFunc func(*mockUser.MockUserRepository, *mockTag.MockTagRepository)
	}

	tests := []testCase{
		{
			name: "存在しないユーザーを新規作成",
			request: &CreateUserRequest{
				Name:     "newUser",
				Email:    "new@example.com",
				Password: "newpassword12345",
				Profile:  "new profile",
				Skills: []SkillRequest{
					{Evaluation: 5, Years: 3},
				},
				Careers: []CareersRequest{
					{Detail: "Software Developer", AdFrom: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), AdTo: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			wantErr: nil,
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository) {
				mockUserRepo.EXPECT().FindByName(gomock.Any(), "newUser").Return(nil, userdm.ErrUserNotFound)
				mockTagRepo.EXPECT().FindByName(gomock.Any(), gomock.Any()).Return(nil, tagdm.ErrTagNotFound)
				mockTagRepo.EXPECT().CreateNewTag(gomock.Any(), gomock.Any()).Return(&tagdm.Tag{}, nil)
				mockUserRepo.EXPECT().Store(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil) // <---- Here

			},
		},
		{
			name: "存在するユーザーを作成",
			request: &CreateUserRequest{
				Name:     "testUser",
				Email:    "test@example.com",
				Password: "password01234",
				Profile:  "test profile",
				Skills: []SkillRequest{
					{Evaluation: 4, Years: 2},
				},
				Careers: []CareersRequest{
					{Detail: "Backend Developer", AdFrom: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), AdTo: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			wantErr: ErrUserNameAlreadyExists,
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository) {
				existingUser := &userdm.User{}
				mockUserRepo.EXPECT().FindByName(gomock.Any(), "testUser").Return(existingUser, nil)
			},
		},
		{
			name: "パスワードが１２文字以下",
			request: &CreateUserRequest{
				Name:     "testUser",
				Email:    "test@example.com",
				Password: "pass",
				Profile:  "test profile",
				Skills: []SkillRequest{
					{Evaluation: 3, Years: 1},
				},
				Careers: []CareersRequest{
					{Detail: "Backend Developer", AdFrom: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), AdTo: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			wantErr: errors.New("userPassword length under 12"),
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository) {
				mockUserRepo.EXPECT().FindByName(gomock.Any(), "testUser").Return(nil, nil)
			},
		},
		{
			name: "新規タグの作成",
			request: &CreateUserRequest{
				Name:     "tagTestUser",
				Email:    "tagtest@example.com",
				Password: "tagpassword12345",
				Profile:  "tag test profile",
				Skills: []SkillRequest{
					{TagName: "Go", Evaluation: 5, Years: 3},
				},
				Careers: nil,
			},
			wantErr: nil,
			mockFunc: func(mockUserRepo *mockUser.MockUserRepository, mockTagRepo *mockTag.MockTagRepository) {
				mockUserRepo.EXPECT().FindByName(gomock.Any(), "tagTestUser").Return(nil, userdm.ErrUserNotFound)
				mockTagRepo.EXPECT().FindByName(gomock.Any(), "Go").Return(nil, tagdm.ErrTagNotFound)
				mockTagRepo.EXPECT().CreateNewTag(gomock.Any(), "Go").Return(&tagdm.Tag{}, nil)
				mockUserRepo.EXPECT().Store(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mockUser.NewMockUserRepository(ctrl)
			mockTagRepo := mockTag.NewMockTagRepository(ctrl)
			tt.mockFunc(mockUserRepo, mockTagRepo)

			service := NewCreateUserAppService(mockUserRepo, mockTagRepo)

			err := service.Exec(context.TODO(), tt.request)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
