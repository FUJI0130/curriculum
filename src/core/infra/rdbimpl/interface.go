package rdbimpl

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBTxRun struct {
	Runner interface {
		Queryx(query string, args ...any) (*sqlx.Rows, error)
		Get(dest any, query string, args ...any) error
		Exec(query string, args ...any) (sql.Result, error)
		Select(dest any, query string, args ...any) error
	}
}

func (r *DBTxRun) Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return r.Runner.Queryx(query, args...)
}

func (r *DBTxRun) Get(dest any, query string, args ...any) error {
	return r.Runner.Get(dest, query, args...)
}

func (r *DBTxRun) Exec(query string, args ...any) (sql.Result, error) {
	return r.Runner.Exec(query, args...)
}

func (r *DBTxRun) Select(dest any, query string, args ...any) error {
	return r.Runner.Select(dest, query, args...)
}
