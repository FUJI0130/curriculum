package userdm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReconstruct(t *testing.T) {
	validTime := time.Now()

	tests := []struct {
		name      string
		id        string
		email     string
		password  string
		profile   string
		createdAt time.Time
		updatedAt time.Time
		wantErr   bool
	}{
		{
			name:      "valid user reconstruction",
			id:        "550e8400-e29b-41d4-a716-446655440000",
			email:     "test@example.com",
			password:  "someValidPassword",
			profile:   "validProfile",
			createdAt: validTime,
			updatedAt: validTime,
			wantErr:   false,
		},
		{
			name:    "invalid email format",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			email:   "invalidEmail",
			wantErr: true,
		},
		// 他のエラーケースやバリデーションケースを追加できます
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := Reconstruct(tt.id, tt.name, tt.email, tt.password, tt.profile, tt.createdAt)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGenWhenCreate(t *testing.T) {

	tests := []struct {
		name     string
		email    string
		password string
		profile  string
		wantErr  bool
	}{
		{
			name:     "valid user generation",
			email:    "test@example.com",
			password: "someValidPassword",
			profile:  "validProfile",
			wantErr:  false,
		},
		{
			name:    "empty profile",
			email:   "test@example.com",
			profile: "",
			wantErr: true,
		},
		// 他のエラーケースやバリデーションケースを追加できます
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := GenWhenCreate("TestName", tt.email, tt.password, tt.profile, []SkillParam{}, []CareerParam{})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTestNewUser(t *testing.T) {

	tests := []struct {
		name    string
		id      string
		email   string
		wantErr bool
	}{
		{
			name:    "valid test user",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "invalid email format",
			id:      "550e8400-e29b-41d4-a716-446655440000",
			email:   "invalidEmail",
			wantErr: true,
		},
		// 他のエラーケースやバリデーションケースを追加できます
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := TestNewUser(tt.id, "TestName", tt.email)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
