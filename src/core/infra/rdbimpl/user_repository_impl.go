package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/tagdm"
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

type SkillRequest struct {
	id         string    `db:"id"`
	tagId      string    `db:"tag_id"`
	userId     string    `db:"user_id"`
	createdAt  time.Time `db:"created_at"`
	updatedAt  time.Time `db:"updated_at"`
	evaluation uint8     `db:"evaluation"`
	years      uint8     `db:"years"`
}

type CareersRequest struct {
	id        string    `db:"id"`
	detail    string    `db:"detail"`
	adFrom    time.Time `db:"ad_from"`
	adTo      time.Time `db:"ad_to"`
	userId    string    `db:"user_id"`
	createdAt time.Time `db:"created_at"`
	updatedAt time.Time `db:"updated_at"`
}

func NewUserRepository(conn *sqlx.DB) userdm.UserRepository {
	return &userRepositoryImpl{Conn: conn}
}

func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User, skills []*userdm.Skills, careers []*userdm.Careers) error {
	queryUser := "INSERT INTO users (id, name, email, password, profile, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	//ここにtagIDの記述を追加
	// tagId, err := tagdm.NewTagID()
	// if err != nil {
	// 	return err
	// }
	_, err := repo.Conn.Exec(queryUser, user.ID(), user.Name(), user.Email(), user.Password(), user.Profile(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())
	if err != nil {
		return err
	}

	// TODO:　後程タグIDの部分は修正する　一時的にこの形で渡すこととする 23/8/30
	tagId, err := tagdm.NewTagID()
	if err != nil {
		return err
	}
	for _, skill := range skills {
		querySkill := "INSERT INTO skills (id,tag_id,user_id,created_at,updated_at, evaluation, years) VALUES (?, ?, ?, ?, ?, ?)"
		_, err = repo.Conn.Exec(querySkill, skill.ID(), tagId.String(), user.ID(), skill.CreatedAt(), skill.UpdatedAt(), skill.Evaluation(), skill.Years())
		if err != nil {
			return err
		}
	}

	for _, career := range careers {
		queryCareer := "INSERT INTO careers (user_id, detail, ad_from, ad_to, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?,?)"
		// ad_from, ad_to を Date() メソッドで取得
		_, err = repo.Conn.Exec(queryCareer, career.ID(), career.Detail(), career.AdFrom(), career.AdTo(), career.UserID(), career.CreatedAt().DateTime(), career.UpdatedAt().DateTime())
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
		log.Println("[DEBUG] FindByName error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Error in FindByName:", err)
			return nil, userdm.ErrUserNotFound
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	user, err := userdm.NewUser(tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt, tempUser.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error converting userRequest to User: %v", err)
	}

	return user, nil
}
