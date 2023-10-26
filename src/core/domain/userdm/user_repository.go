//go:generate mockgen -destination=../../mock/mock_user/user_repository_mock.go -package=mock_userdm . UserRepository

package userdm

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, userdomain *UserDomain) error
	FindByUserName(ctx context.Context, name string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserNames(ctx context.Context, names []string) (map[string]*User, error)
	FindSkillsByUserID(ctx context.Context, userID string) ([]Skill, error)   // 新しいメソッド
	FindCareersByUserID(ctx context.Context, userID string) ([]Career, error) // 新しいメソッド
	UpdateUser(ctx context.Context, user *User) error
	UpdateSkill(ctx context.Context, skill *Skill) error    // 新しいメソッド
	UpdateCareer(ctx context.Context, career *Career) error // 新しいメソッド
}
