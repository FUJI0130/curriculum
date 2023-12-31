package tagdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type TagID string

func NewTagID() (TagID, error) {
	tagID, err := uuid.NewRandom()
	if err != nil {
		return TagID(""), customerrors.WrapUnprocessableEntityError(err, "ID is error")
	}
	return TagID(tagID.String()), nil
}
func NewTagIDFromString(idStr string) (TagID, error) {
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.WrapUnprocessableEntityError(err, "ID is error")
	}
	return TagID(idStr), nil
}
func (id TagID) String() string {
	return string(id)
}

func (tagID1 TagID) Equal(tagID2 TagID) bool {
	return tagID1 == tagID2
}
