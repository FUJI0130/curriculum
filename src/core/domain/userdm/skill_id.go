package userdm

import (
	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
	"github.com/google/uuid"
)

type SkillID string

func NewSkillID() (SkillID, error) {
	skillID, err := uuid.NewRandom()
	if err != nil {
		return SkillID(""), customerrors.ErrInvalidSkillIDFormat(err, "NewSkillID")
	}
	skillIDValueObject := SkillID(skillID.String())
	return skillIDValueObject, nil
}
func NewSkillIDFromString(idStr string) (SkillID, error) {
	// UUIDの形式であるか確認
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.ErrInvalidSkillIDFormat(err, "NewSkillIDFromString")
	}
	return SkillID(idStr), nil
}
func (id SkillID) String() string {
	return string(id)
}

func (tagID1 SkillID) Equal(tagID2 SkillID) bool {
	return tagID1 == tagID2
}
