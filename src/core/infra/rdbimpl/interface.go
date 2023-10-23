package rdbimpl

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Queryer interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
}
