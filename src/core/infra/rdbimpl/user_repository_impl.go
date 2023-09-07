package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	Conn *sqlx.DB
}

type userRequest struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Profile   string    `db:"profile"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserRepository(conn *sqlx.DB) userdm.UserRepository {
	return &userRepositoryImpl{Conn: conn}
}

// error　が戻り値の型
func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User, skill []*userdm.Skill, career []*userdm.Career) error {

	queryUser := "INSERT INTO users (id, name, email, password, profile, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.Conn.Exec(queryUser, user.ID().String(), user.Name(), user.Email(), user.Password(), user.Profile(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())
	if err != nil {
		return err
	}

	// Skillsの保存
	for _, skill := range skill {
		querySkill := "INSERT INTO skills (id,tag_id,user_id,created_at,updated_at, evaluation, years) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = repo.Conn.Exec(querySkill, skill.ID().String(), skill.TagID().String(), user.ID().String(), skill.CreatedAt().DateTime(), skill.UpdatedAt().DateTime(), skill.Evaluation().Value(), skill.Year().Value())
		if err != nil {
			return err
		}
	}

	// Careersの保存

	for _, career := range career {
		queryCareer := "INSERT INTO careers (id,user_id, detail, ad_from, ad_to, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
		// ad_from, ad_to を Date() メソッドで取得
		_, err = repo.Conn.Exec(queryCareer, career.ID().String(), career.UserID().String(), career.Detail(), career.AdFrom(), career.AdTo(), career.CreatedAt().DateTime(), career.UpdatedAt().DateTime())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *userRepositoryImpl) FindByName(ctx context.Context, name string) (*userdm.User, error) {
	query := "SELECT * FROM users WHERE name = ?"

	var tempUser userRequest
	err := repo.Conn.Get(&tempUser, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, userdm.ErrUserNotFound
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	// userRequestからuserdm.Userへの変換
	userID := userdm.UserID(tempUser.ID) // UserID の型に合わせて変換が必要な場合
	user, err := userdm.ReconstructUser(userID, tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt, tempUser.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error reconstructing user from userRequest: %v", err)
	}

	return user, nil
}
