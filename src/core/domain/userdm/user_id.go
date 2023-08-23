package userdm

import (
	"fmt"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() (UserID, error) {
	userId, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return UserID(""), err
	}
	userId_VO := UserID(userId.String())
	return userId_VO, nil
}

func (id UserID) String() string {
	uuidString := string(id)
	return uuidString
}

func (userID1 UserID) Equal(userID2 UserID) bool {
	return userID1 == userID2
}
