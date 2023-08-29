package userdm

type Skills struct {
	evaluation SkillsEvaluation
	years      SkillsYears
}

func NewSkills(evaluation int, years int) (*Skills, error) {
	eval, err := NewSkillsEvaluation(evaluation)
	if err != nil {
		return nil, err
	}

	y, err := NewSkillsYears(years)
	if err != nil {
		return nil, err
	}

	return &Skills{
		evaluation: *eval,
		years:      *y,
	}, nil
}

func (s *Skills) Evaluation() SkillsEvaluation {
	return s.evaluation
}

func (s *Skills) Years() SkillsYears {
	return s.years
}
