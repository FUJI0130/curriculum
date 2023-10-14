package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type CareerID string

func NewCareerID() (CareerID, error) {
	careerID, err := uuid.NewRandom()
	if err != nil {
		return CareerID(""), customerrors.WrapUnprocessableEntityError(err, "ID is error")
	}
	return CareerID(careerID.String()), nil
}

func (id CareerID) String() string {
	return string(id)
}

func (careerID1 CareerID) Equal(careerID2 CareerID) bool {
	return careerID1 == careerID2
}
