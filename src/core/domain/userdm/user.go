package userdm

import (
	"time"
	"unicode/utf8"

	"github.com/FUJI0130/curriculum/src/core/domain/shared/sharedvo"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

const nameMaxlength = 256
const profileMaxlength = 256

type User struct {
	id        UserID             `db:"id"`
	name      string             `db:"name"`
	email     UserEmail          `db:"email"`
	password  UserPassword       `db:"password"`
	profile   string             `db:"profile"`
	skills    []Skill            `db:"skills"`
	careers   []Career           `db:"careers"`
	createdAt sharedvo.CreatedAt `db:"created_at"`
	updatedAt sharedvo.UpdatedAt `db:"updated_at"`
}

func NewUser(name string, email string, password string, profile string, skills []Skill, careers []Career) (*User, error) {

	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("Username is empty")
	}

	userID, err := NewUserID()
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
		return nil, customerrors.NewUnprocessableEntityError("UserProfile is empty")
	} else if profileMaxlength < countProfile {
		return nil, customerrors.NewUnprocessableEntityError("UserProfile is too long")
	}

	userCreatedAt := sharedvo.NewCreatedAt()
	userUpdatedAt := sharedvo.NewUpdatedAt()

	return &User{
		id:        userID,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		skills:    skills,
		careers:   careers,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func TestNewUser(id string, name string, email string) (*User, error) {
	userID, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}

	userEmail, err := NewUserEmail(email)
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
		return nil, customerrors.NewUnprocessableEntityError("profile is empty")
	} else if profileMaxlength < countProfile {
		return nil, customerrors.NewUnprocessableEntityError("profile is too long")
	}
	userCreatedAt := sharedvo.NewCreatedAt()
	userUpdatedAt := sharedvo.NewUpdatedAt()

	return &User{
		id:        userID,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   userProfile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func ReconstructEntity(id string, name string, email string, password string, profile string, skills []Skill, careers []Career, createdAt time.Time) (*User, error) {
	userID, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
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
	if err != nil {
		return nil, err
	}

	return &User{
		id:        userID,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   profile,
		skills:    skills,
		careers:   careers,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func ReconstructUser(id string, name string, email string, password string, profile string, createdAt time.Time) (*User, error) {
	userID, err := NewUserIDFromString(id)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return nil, customerrors.NewUnprocessableEntityError("name is empty")
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
	if err != nil {
		return nil, err
	}
	return &User{
		id:        userID,
		name:      name,
		email:     userEmail,
		password:  userPassword,
		profile:   profile,
		createdAt: userCreatedAt,
		updatedAt: userUpdatedAt,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() UserEmail {
	return u.email
}

func (u *User) Password() UserPassword {
	return u.password
}

func (u *User) Profile() string {
	return u.profile
}
func (u *User) Skills() []Skill {
	return u.skills
}

func (u *User) Careers() []Career {
	return u.careers
}

func (u *User) CreatedAt() sharedvo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() sharedvo.UpdatedAt {
	return u.updatedAt
}

func (u *User) Update(name string, email string, password string, profile string, skills []Skill, careers []Career) error {
	// 名前の検証
	if name == "" {
		return customerrors.NewUnprocessableEntityError("name is empty")
	}
	if utf8.RuneCountInString(name) > nameMaxlength {
		return customerrors.NewUnprocessableEntityError("name is too long")
	}

	// EmailとPasswordの値オブジェクトの生成
	userEmail, err := NewUserEmail(email)
	if err != nil {
		return err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return err
	}

	// 更新処理
	u.name = name
	u.email = userEmail
	u.password = userPassword

	// プロフィールの検証
	if profile == "" {
		return customerrors.NewUnprocessableEntityError("UserProfile is empty")
	}
	if utf8.RuneCountInString(profile) > profileMaxlength {
		return customerrors.NewUnprocessableEntityError("UserProfile is too long")
	}

	u.profile = profile
	u.skills = skills   // スキルのスライスを更新
	u.careers = careers // キャリアのスライスを更新
	u.updatedAt = sharedvo.NewUpdatedAt()
	if err != nil {
		return err
	}

	return nil
}
