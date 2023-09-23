package tagdm

import (
	"errors"

	"github.com/google/uuid"
)

type TagID string

func NewTagID() (TagID, error) {
	tagID, err := uuid.NewRandom()
	if err != nil {
		return TagID(""), err
	}
	tagIDValueObject := TagID(tagID.String())
	return tagIDValueObject, nil
}
func NewTagIDFromString(idStr string) (TagID, error) {
	// UUIDの形式であるか確認
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", errors.New("invalid UUID format")
	}
	return TagID(idStr), nil
}
func (id TagID) String() string {
	uuidString := string(id)
	return uuidString
}

func (tagID1 TagID) Equal(tagID2 TagID) bool {
	return tagID1 == tagID2
}
