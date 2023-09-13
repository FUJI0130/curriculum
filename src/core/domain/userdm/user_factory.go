package userdm

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
)

type UserFactory interface {
	Reconstruct(id string, name string, email string, password string, profile string, createdAt time.Time) (*User, error)
	GenWhenCreate(name string, email string, password string, profile string) (*User, error)
	TestNewUser(id string, name string, email string) (*User, error)
}

type userFactory struct{}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func (f *userFactory) Reconstruct(id string, name string, email string, password string, profile string, createdAt time.Time) (*User, error) {
	userId, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}

	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}
	userCreatedAt, err := sharedvo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}

	userUpdatedAt, err := sharedvo.NewUpdatedAtByVal(time.Now())
	fmt.Println("Reconstruct UpdatedAt:", userUpdatedAt.DateTime())
	if err != nil {
		return nil, err
	}
	return &User{
		id:        userId,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   profile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func (f *userFactory) GenWhenCreate(name string, email string, password string, profile string) (*User, error) {
	userId, err := NewUserID()
	if err != nil {
		return nil, err
	}

	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}
	userProfile := profile
	countProfile := utf8.RuneCountInString(profile)
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if nameMaxlength < countProfile {
		return nil, errors.New("userProfile length over nameMaxlength")
	}

	userCreatedAt := sharedvo.NewCreatedAt()
	userUpdatedAt := sharedvo.NewUpdatedAt()

	return &User{
		id:        userId,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func (f *userFactory) TestNewUser(id string, name string, email string) (*User, error) {
	userId, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}

	userEmail, err := NewUserEmail(email) // ここでバリデーションを省略しています。必要に応じて調整してください。
	if err != nil {
		return nil, err
	}
	password := "newuserpassword"
	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}
	userProfile := "親譲りの無鉄砲で小供の時から損ばかりしている。小学校に居る時分学校の二階から飛び降りて一週間ほど腰を抜かした事がある。なぜそんな無闇をしたと聞く人があるかも知れぬ。別段深い理由でもない。新築の二階から首を出していたら、同級生の一人が冗談に、いくら威張っても、そこから飛び降りる事は出来まい。弱虫やーい。と囃したからである。小使に負ぶさって帰って来た時、おやじが大きな眼をして二階ぐらいから飛び降りて腰を抜かす奴があるかと云ったから、この次は抜かさずに飛んで見せますと答えた。（青空文庫より）親譲りの無鉄砲で小供"
	countProfile := utf8.RuneCountInString(userProfile)
	if userProfile == "" {
		return nil, errors.New("userProfile cannot be empty")
	} else if nameMaxlength < countProfile {
		return nil, errors.New("userProfile length over nameMaxlength")
	}
	userCreatedAt := sharedvo.NewCreatedAt()
	userUpdatedAt := sharedvo.NewUpdatedAt()

	return &User{
		id:        userId,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}
