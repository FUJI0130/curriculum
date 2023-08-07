package userdm

import "errors"

type UserProfile string

func NewUserProfile(userProfile string) (*UserProfile, error) {
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	}
	userProfile_VO := UserProfile(userProfile)
	return &userProfile_VO, nil
}

func (profile *UserProfile) String() string {
	return string(*profile)
}

func (userProfile1 *UserProfile) Equal(userProfile2 *UserProfile) bool {
	return *userProfile1 == *userProfile2
}
