package userdm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {

	createdAt := time.Now()
	updatedAt := time.Now()

	user, err := NewUser("TestName", "test@example.com", "Password", "Profile", createdAt, updatedAt)

	// assert.NoError(t, err)
	require.NoError(t, err)
	assert.Equal(t, "TestName", user.name)
	assert.Equal(t, "test@example.com", user.email)
	assert.Equal(t, "Password", user.password)
	assert.Equal(t, "Profile", user.profile)
	assert.Equal(t, createdAt, user.createdAt)
	assert.Equal(t, updatedAt, user.updatedAt)
}

func TestUserID_Equal(t *testing.T) {
	userID1, _ := NewUserID()
	userID2, _ := NewUserID()

	assert.False(t, userID1.Equal(userID2))

	userID3, _ := NewUserID()
	assert.True(t, userID1.Equal(userID3))
}
