package userdm

import (
	"github.com/google/uuid"
)

type CareerID string

func NewCareerID() (CareerID, error) {
	careerID, err := uuid.NewRandom()
	if err != nil {
		return CareerID(""), err
	}
	careerIDValueObject := CareerID(careerID.String())
	return careerIDValueObject, nil
}

func (id CareerID) String() string {
	uuidString := string(id)
	return uuidString
}

func (careerID1 CareerID) Equal(careerID2 CareerID) bool {
	return careerID1 == careerID2
}
