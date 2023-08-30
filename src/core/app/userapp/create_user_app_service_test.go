package userapp

import (
	"context"
	"errors"
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAppService_Exec(t *testing.T) {
	type testCase struct {
		name     string
		request  *CreateUserRequest
		wantErr  error
		mockFunc func(*mockUser.MockUserRepository)
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
				Careers: []userdm.CareersRequest{
					{From: 2020, To: 2023, Detail: "Software Developer"},
				},
			},
			wantErr: nil,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(gomock.Any(), "newUser").Return(nil, userdm.ErrUserNotFound)
				mockRepo.EXPECT().Store(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
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
				Careers: []userdm.CareersRequest{
					{From: 2021, To: 2023, Detail: "Backend Developer"},
				},
			},
			wantErr: ErrUserNameAlreadyExists,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				existingUser := &userdm.User{}
				mockRepo.EXPECT().FindByName(gomock.Any(), "testUser").Return(existingUser, nil)
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
				Careers: []userdm.CareersRequest{
					{From: 2022, To: 2023, Detail: "Frontend Developer"},
				},
			},
			wantErr: errors.New("userPassword length under 12"),
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(gomock.Any(), "testUser").Return(nil, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockUser.NewMockUserRepository(ctrl)
			tt.mockFunc(mockRepo)

			service := NewCreateUserAppService(mockRepo)

			err := service.Exec(context.TODO(), tt.request)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
