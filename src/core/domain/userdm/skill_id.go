package userdm

import (
	"errors"

	"github.com/google/uuid"
)

type SkillID string

func NewSkillID() (SkillID, error) {
	skillID, err := uuid.NewRandom()
	if err != nil {
		return SkillID(""), err
	}
	skillIDValueObject := SkillID(skillID.String())
	return skillIDValueObject, nil
}
func NewsSkillIDFromString(idStr string) (SkillID, error) {
	// UUIDの形式であるか確認
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", errors.New("invalid UUID format")
	}
	return SkillID(idStr), nil
}
func (id SkillID) String() string {
	uuidString := string(id)
	return uuidString
}

func (tagID1 SkillID) Equal(tagID2 SkillID) bool {
	return tagID1 == tagID2
}
