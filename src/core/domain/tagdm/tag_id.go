package tagdm

import (
	"fmt"

	"github.com/google/uuid"
)

type TagID string

func NewTagID() (TagID, error) {
	tagId, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return TagID(""), err
	}
	tagId_VO := TagID(tagId.String())
	return tagId_VO, nil
}

func (id TagID) String() string {
	uuidString := string(id)
	return uuidString
}

func (tagID1 TagID) Equal(tagID2 TagID) bool {
	return tagID1 == tagID2
}
