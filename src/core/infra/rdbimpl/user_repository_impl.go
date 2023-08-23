// - ORMを使ってDBの操作の実実装を行う
// - FindByName、Storeの処理をそれぞれORMを使って実装する

package rdbimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	Conn *sqlx.DB
}

type scanUser struct {
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

func (repo *userRepositoryImpl) Store(ctx context.Context, user *userdm.User) error {
	query := "INSERT INTO users (id, name, email, password, profile, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.Conn.Exec(query, user.ID(), user.Name(), user.Email(), user.Password(), user.Profile(), user.CreatedAt().DateTime(), user.UpdatedAt().DateTime())

	return err
}

func (repo *userRepositoryImpl) FindByName(ctx context.Context, name string) (*userdm.User, error) {
	query := "SELECT * FROM users WHERE name = ?"
	var tempUser scanUser
	err := repo.Conn.Get(&tempUser, query, name)
	if err != nil {
		log.Println("[DEBUG] FindByName error:", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Error in FindByName:", err)
			return nil, userdm.ErrUserNotFound
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	// scanUserをuserdm.Userに変換
	user, err := userdm.NewUser(tempUser.Name, tempUser.Email, tempUser.Password, tempUser.Profile, tempUser.CreatedAt, tempUser.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error converting scanUser to User: %v", err)
	}

	return user, nil
}
