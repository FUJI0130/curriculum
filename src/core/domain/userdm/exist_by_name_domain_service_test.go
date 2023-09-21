// package interfaces
// package userdm_test
package userdm

import (
	"context"
	"errors"
	"testing"

	// "github.com/FUJI0130/curriculum/src/core/domain/userdm"
	// "github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	// "github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/FUJI0130/curriculum/src/core/mock/mockUser"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExistByNameDomainService_IsExist(t *testing.T) {
	mockName := "Test User"

	ctx := context.TODO()

	tests := []struct {
		name      string
		inputName string
		// mockFunc   func(mockRepo *mockUser.MockUserRepository)
		mockFunc   func(mockRepo *mockUser.MockUserRepository)
		wantResult bool
		wantErr    error
	}{
		{
			name:      "User exists",
			inputName: mockName,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(ctx, mockName).Return(&User{}, nil).Times(1)
			},
			wantResult: true,
			wantErr:    nil,
		},
		{
			name:      "User does not exist",
			inputName: mockName,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(ctx, mockName).Return(nil, ErrUserNotFound).Times(1)
			},
			wantResult: false,
			wantErr:    nil,
		},
		{
			name:      "Repository error",
			inputName: mockName,
			mockFunc: func(mockRepo *mockUser.MockUserRepository) {
				mockRepo.EXPECT().FindByName(ctx, mockName).Return(nil, errors.New("unexpected error")).Times(1)
			},
			wantResult: false,
			wantErr:    errors.New("unexpected error"),
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mockUser.NewMockUserRepository(ctrl)

			domainService := NewExistByNameDomainService(mockRepo)

			tt.mockFunc(mockRepo)

			result, err := domainService.IsExist(ctx, tt.inputName)
			assert.Equal(t, tt.wantResult, result)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
