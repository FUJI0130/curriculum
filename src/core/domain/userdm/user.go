package userdm

import (
	"errors"
	"time"
)

type User struct {
	id        UserID
	name      UserName
	email     UserEmail
	password  UserPassword
	profile   UserProfile
	createdAt UserCreatedAt
	updatedAt UserUpdatedAt
}

func NewUser(name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if name == "" || password == "" {
		return nil, errors.New("name and password cannot be empty")
	}
	user_id_VO, _ := NewUserID()
	user_name_VO, _ := NewUserName(name)
	user_email_VO, _ := NewUserEmail(email)
	user_password_VO, _ := NewUserPassword(password)
	user_profile_VO, _ := NewUserProfile(profile)
	user_createdAt_VO, _ := NewUserCreatedAt(createdAt)
	user_updatedAt_VO, _ := NewUserUpdatedAt(updatedAt)

	return &User{
		id:        *user_id_VO,
		name:      *user_name_VO,
		email:     *user_email_VO,
		password:  *user_password_VO,
		profile:   *user_profile_VO,
		createdAt: *user_createdAt_VO,
		updatedAt: *user_updatedAt_VO,
	}, nil
}
