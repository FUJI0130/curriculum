package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type UserID string

func NewUserID() (UserID, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return UserID(""), customerrors.WrapUnprocessableEntityError(err, "ID is error")
	}
	userIDValueObject := UserID(userID.String())
	return userIDValueObject, nil
}

func NewUserIDFromString(idStr string) (UserID, error) {
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.WrapUnprocessableEntityError(err, "ID is error")
	}
	return UserID(idStr), nil
}
func (id UserID) String() string {
	return string(id)
}

func (userID1 UserID) Equal(userID2 UserID) bool {
	return userID1 == userID2
}
