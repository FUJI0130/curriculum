package userdm

import (
	"fmt"

	"github.com/google/uuid"
)

type CareerID string

func NewCareerID() (CareerID, error) {
	careerId, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return CareerID(""), err
	}
	careerId_VO := CareerID(careerId.String())
	return careerId_VO, nil
}

func (id CareerID) String() string {
	uuidString := string(id)
	return uuidString
}

func (careerID1 CareerID) Equal(careerID2 CareerID) bool {
	return careerID1 == careerID2
}
