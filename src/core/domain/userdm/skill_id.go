package userdm

import (
	"fmt"

	"github.com/google/uuid"
)

type SkillID string

func NewSkillID() (SkillID, error) {
	skillId, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return SkillID(""), err
	}
	skillId_VO := SkillID(skillId.String())
	return skillId_VO, nil
}

func (id SkillID) String() string {
	uuidString := string(id)
	return uuidString
}

func (tagID1 SkillID) Equal(tagID2 SkillID) bool {
	return tagID1 == tagID2
}
