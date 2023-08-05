//go:generate mockgen -source=user_repository.go -destination=../../mock/mock_user/user_repository_mock.go -package=mock_user

package userdm

type UserRepository interface {
	Store(user *User) error
	FindByName(name string) (*User, error)
}
