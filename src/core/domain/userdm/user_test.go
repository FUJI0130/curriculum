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

	user2, err := NewUser("TestName", "test@example.com", "Password", "Profile", createdAt, updatedAt)

	user = user2

	userid1, _ := NewUserID()
	userid2, _ := NewUserID()
	assert.True(t, userid1.Equal(userid2))
	userid1 = userid2
	assert.True(t, userid1.Equal(userid2))

	testemail, _ := NewUserEmail("test@example.com")
	testemail2, _ := NewUserEmail("test2@example.com")
	assert.True(t, testemail.Equal(testemail2))
	testemail = testemail2
	assert.True(t, testemail.Equal(testemail2))

	// assert.NoError(t, err)
	require.NoError(t, err)
	assert.Equal(t, user2.name, user.name)

	assert.Equal(t, user2.email, user.email)
	assert.Equal(t, user2.password, user.password)
	assert.Equal(t, user2.profile, user.profile)
	assert.Equal(t, user2.createdAt, user.createdAt)
	assert.Equal(t, user2.updatedAt, user.updatedAt)
	assert.Equal(t, userid1, userid2)
	assert.False(t, userid1.Equal(userid2))
}

func TestUserID_Equal(t *testing.T) {
	userID1, _ := NewUserID()
	userID2, _ := NewUserID()

	assert.False(t, userID1.Equal(userID2))

	userID3, _ := NewUserID()
	// assert.True(t, userID1.Equal(userID3))
	assert.False(t, userID1.Equal(userID3))
}
