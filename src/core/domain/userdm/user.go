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
	user_id, err := NewUserID()
	if err != nil {
		return nil, err
	}

	user_name, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	user_email, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	user_password, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	user_profile, err := NewUserProfile(profile)
	if err != nil {
		return nil, err
	}

	user_createdAt, err := NewUserCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}

	user_updatedAt, err := NewUserUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        user_id,
		name:      *user_name,
		email:     *user_email,
		password:  *user_password,
		profile:   *user_profile,
		createdAt: *user_createdAt,
		updatedAt: *user_updatedAt,
	}, nil
}
