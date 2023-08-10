package userdm

import "errors"

type UserProfile string

func NewUserProfile(userProfile string) (*UserProfile, error) {
	count := len([]rune(userProfile))
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if 255 < count {
		return nil, errors.New("userProfile length over 256")
	}
	userProfile_VO := UserProfile(userProfile)
	return &userProfile_VO, nil
}

func (profile *UserProfile) String() string {
	userProfileString := string(*profile)
	return userProfileString
}

func (userProfile1 *UserProfile) Equal(userProfile2 *UserProfile) bool {
	return *userProfile1 == *userProfile2
}
