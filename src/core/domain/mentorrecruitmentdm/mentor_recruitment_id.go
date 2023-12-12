package mentorrecruitmentdm

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/google/uuid"
)

type MentorRecruitmentID string

func NewMentorRecruitmentID() (MentorRecruitmentID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return MentorRecruitmentID(""), customerrors.WrapUnprocessableEntityError(err, "ID generation error")
	}
	return MentorRecruitmentID(id.String()), nil
}

func NewMentorRecruitmentIDFromString(idStr string) (MentorRecruitmentID, error) {
	_, err := uuid.Parse(idStr)
	if err != nil {
		return "", customerrors.WrapUnprocessableEntityError(err, "ID parsing error")
	}
	return MentorRecruitmentID(idStr), nil
}

func (id MentorRecruitmentID) String() string {
	return string(id)
}

func (id1 MentorRecruitmentID) Equal(id2 MentorRecruitmentID) bool {
	return id1 == id2
}
