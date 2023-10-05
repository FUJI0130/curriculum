package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type UserID string

func NewUserID() (UserID, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return UserID(""), customerrors.WrapUnprocessableEntityError(err, "NewUserID")
	}
	userIDValueObject := UserID(userID.String())
	return userIDValueObject, nil
}

func NewUserIDFromString(idStr string) (UserID, error) {
	// UUIDの形式であるか確認
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.WrapUnprocessableEntityError(err, "NewUserIDFromString")
	}
	return UserID(idStr), nil
}
func (id UserID) String() string {
	return string(id)
}

func (userID1 UserID) Equal(userID2 UserID) bool {
	return userID1 == userID2
}
