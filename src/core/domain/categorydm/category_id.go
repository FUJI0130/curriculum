package categorydm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type CategoryID string

func NewCategoryID() (CategoryID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return CategoryID(""), customerrors.WrapUnprocessableEntityError(err, "ID generation error")
	}
	return CategoryID(id.String()), nil
}

func NewCategoryIDFromString(idStr string) (CategoryID, error) {
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.WrapUnprocessableEntityError(err, "ID parsing error")
	}
	return CategoryID(idStr), nil
}

func (id CategoryID) String() string {
	return string(id)
}

func (id1 CategoryID) Equal(id2 CategoryID) bool {
	return id1 == id2
}
