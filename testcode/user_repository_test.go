package testcode

import (
	"testing"
	"time"

	"github.com/FUJI0130/curriculum/src/core/mock/mock_user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	userID, _ := NewUserID("test-uuid")
	createdAt := time.Now()
	updatedAt := time.Now()
	user, _ := NewUser(userID, "TestName", "test@example.com", "Password", "Profile", createdAt, updatedAt)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックの設定
	repo := mock_user.NewMockUserRepository(ctrl)
	repo.EXPECT().FindByName("TestName").Return(user, nil).Times(1)
	repo.EXPECT().Store(user).Return(nil).Times(1)

	// FindByNameのテスト
	foundUser, err := repo.FindByName("TestName")
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	// Storeのテスト
	err = repo.Store(user)
	assert.NoError(t, err)
}
