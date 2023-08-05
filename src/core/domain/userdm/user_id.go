package userdm

import "errors"

type UserID struct {
	value string
}

func NewUserID(userID string) (*UserID, error) {
	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}
	return &UserID{value: userID}, nil
}

func (id *UserID) String() string {
	return id.value
}

func (id *UserID) Equal(userID2 *UserID) bool {
	return id.value == userID2.value
}
