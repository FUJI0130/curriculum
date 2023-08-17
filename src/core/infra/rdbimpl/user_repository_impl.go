// - ORMを使ってDBの操作の実実装を行う
// - FindByName、Storeの処理をそれぞれORMを使って実装する

package rdbimpl

import (
	"context"
	"your_path/core/domain/userdm"
	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	Conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) userdm.UserRepository {
	return &userRepositoryImpl{Conn: conn}
}

func (repo *userRepositoryImpl) Store(user *userdm.User) error {
	query := "INSERT INTO users (id, name, password, introduction) VALUES (?, ?, ?, ?)"
	_, err := repo.Conn.Exec(query, user.ID.String(), user.Name, user.Password, user.Introduction)
	return err
}

func (repo *userRepositoryImpl) FindByName(name string) (*userdm.User, error) {
	query := "SELECT * FROM users WHERE name = ?"
	var user userdm.User
	err := repo.Conn.Get(&user, query, name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
