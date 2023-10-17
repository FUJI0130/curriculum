package userdm_test

import (
	"context"
	"testing"

	mockExistByNameDomainService "github.com/FUJI0130/curriculum/src/core/mock/mock_exist_by_name_domain_service"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExistByNameDomainService_Exec(t *testing.T) {
	mockName := "Test User"
	ctx := context.TODO()

	tests := []struct {
		title      string
		inputName  string
		mockFunc   func(mockService *mockExistByNameDomainService.MockExistByNameDomainService)
		wantResult bool
		wantErr    error
	}{
		{
			title:     "User exists",
			inputName: mockName,
			mockFunc: func(mockService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockService.EXPECT().Exec(ctx, mockName).Return(true, nil).Times(1)
			},
			wantResult: true,
			wantErr:    nil,
		},
		{
			title:     "User does not exist",
			inputName: mockName,
			mockFunc: func(mockService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockService.EXPECT().Exec(ctx, mockName).Return(false, nil).Times(1)
			},
			wantResult: false,
			wantErr:    nil,
		},
		{
			title:     "Service error",
			inputName: mockName,
			mockFunc: func(mockService *mockExistByNameDomainService.MockExistByNameDomainService) {
				mockService.EXPECT().Exec(ctx, mockName).Return(false, customerrors.NewNotFoundError("test Service error")).Times(1)
			},
			wantResult: false,
			wantErr:    customerrors.NewNotFoundError("test Service error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := mockExistByNameDomainService.NewMockExistByNameDomainService(ctrl)

			tt.mockFunc(mockService)

			result, err := mockService.Exec(ctx, tt.inputName)
			assert.Equal(t, tt.wantResult, result)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
