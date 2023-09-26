package userdm

type UserDomain struct {
	User    *User
	Skills  []*Skill
	Careers []*Career
}

func NewUserDomain(user *User, skills []*Skill, careers []*Career) *UserDomain {
	return &UserDomain{
		User:    user,
		Skills:  skills,
		Careers: careers,
	}
}
func (ud *UserDomain) Reconstruct(user *User, skills []*Skill, careers []*Career) *UserDomain {
	return &UserDomain{
		User:    user,
		Skills:  skills,
		Careers: careers,
	}
}

func (ud *UserDomain) AddSkill(skill *Skill) {
	ud.Skills = append(ud.Skills, skill)
}

func (ud *UserDomain) AddCareer(career *Career) {
	ud.Careers = append(ud.Careers, career)
}

func (ud *UserDomain) RemoveSkill(skillID SkillID) {
	var newSkills []*Skill
	for _, s := range ud.Skills {
		if s.ID() != skillID {
			newSkills = append(newSkills, s)
		}
	}
	ud.Skills = newSkills
}

func (ud *UserDomain) RemoveCareer(careerID CareerID) {
	var newCareers []*Career
	for _, c := range ud.Careers {
		if c.ID() != careerID {
			newCareers = append(newCareers, c)
		}
	}
	ud.Careers = newCareers
}
