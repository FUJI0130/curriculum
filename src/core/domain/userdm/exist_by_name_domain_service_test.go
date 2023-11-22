package userdm_test

import (
	"context"
	"errors"
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
			title:     "Service internal error",
			inputName: mockName,
			mockFunc: func(mockService *mockExistByNameDomainService.MockExistByNameDomainService) {
				internalError := errors.New("internal error")
				wrappedError := customerrors.WrapInternalServerError(internalError, "An internal error occurred")
				mockService.EXPECT().Exec(ctx, mockName).Return(false, wrappedError).Times(1)
			},
			wantResult: false,
			wantErr:    customerrors.WrapInternalServerError(errors.New("internal error"), "An internal error occurred"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mockExistByNameDomainService.NewMockExistByNameDomainService(ctrl)

			tt.mockFunc(mockService)

			result, err := mockService.Exec(ctx, tt.inputName)
			assert.Equal(t, tt.wantResult, result)
			if tt.wantErr != nil {
				var internalErr *customerrors.InternalServerErrorType
				if errors.As(err, &internalErr) {
					assert.Contains(t, internalErr.Error(), "An internal error occurred")
				} else {
					t.Errorf("Expected an internal server error, got %v", err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
