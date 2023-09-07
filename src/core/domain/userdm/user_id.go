package userdm

import (
	"fmt"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() (UserID, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return UserID(""), err
	}
	userIDValueObject := UserID(userID.String())
	return userIDValueObject, nil
}

func (id UserID) String() string {
	uuidString := string(id)
	return uuidString
}

func (userID1 UserID) Equal(userID2 UserID) bool {
	return userID1 == userID2
}
