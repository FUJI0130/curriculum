package userdm

import (
	"errors"
	"time"
)

type User struct {
	ID        UserID
	Name      string
	Email     string
	Password  string
	Profile   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id *UserID, name string, email string, password string, profile string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if name == "" || password == "" {
		return nil, errors.New("name and password cannot be empty")
	}

	return &User{
		ID:        *id,
		Name:      name,
		Email:     email,
		Password:  password,
		Profile:   profile,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
