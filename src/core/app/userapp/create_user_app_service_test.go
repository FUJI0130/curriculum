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
			},
			wantErr: nil,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				// この名前のユーザーはまだ存在しないことを模倣する
				mockRepo.EXPECT().FindByName(gomock.Any(), "newUser").Return(nil, userdm.ErrUserNotFound)

				// 新しいユーザーの保存を模倣する
				mockRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "存在するユーザーを作成",
			request: &CreateUserRequest{
				Name:     "testUser",
				Email:    "test@example.com",
				Password: "password01234",
				Profile:  "test profile",
			},
			wantErr: nil,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(gomock.Any(), "testUser").Return(nil, nil)
				mockRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "パスワードが１２文字以下",
			request: &CreateUserRequest{
				Name:     "testUser",
				Email:    "test@example.com",
				Password: "pass",
				Profile:  "test profile",
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
