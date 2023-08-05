package userdm

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	userID, _ := NewUserID(uuid.New().String())
	createdAt := time.Now()
	updatedAt := time.Now()

	user, err := NewUser(userID, "TestName", "test@example.com", "Password", "Profile", createdAt, updatedAt)

	assert.NoError(t, err)
	assert.Equal(t, "TestName", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "Password", user.Password)
	assert.Equal(t, "Profile", user.Profile)
	assert.Equal(t, createdAt, user.CreatedAt)
	assert.Equal(t, updatedAt, user.UpdatedAt)
}

func TestUserID_Equal(t *testing.T) {
	userID1, _ := NewUserID(uuid.New().String())
	userID2, _ := NewUserID(uuid.New().String())

	assert.False(t, userID1.Equal(userID2))

	userID3, _ := NewUserID(userID1.String())
	assert.True(t, userID1.Equal(userID3))
}
