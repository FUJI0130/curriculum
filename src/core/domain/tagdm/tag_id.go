package tagdm

import (
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

func (id TagID) String() string {
	uuidString := string(id)
	return uuidString
}

func (tagID1 TagID) Equal(tagID2 TagID) bool {
	return tagID1 == tagID2
}
