package rdbimpl

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBTxRun struct {
	DBoperator dbOperator
}

type dbOperator interface {
	Queryx(query string, args ...any) (*sqlx.Rows, error)
	Get(dest any, query string, args ...any) error
	Exec(query string, args ...any) (sql.Result, error)
	Select(dest any, query string, args ...any) error
}

func (r *DBTxRun) Queryx(query string, args ...any) (*sqlx.Rows, error) {
	return r.DBoperator.Queryx(query, args...)
}

func (r *DBTxRun) Get(dest any, query string, args ...any) error {
	return r.DBoperator.Get(dest, query, args...)
}

func (r *DBTxRun) Exec(query string, args ...any) (sql.Result, error) {
	return r.DBoperator.Exec(query, args...)
}

func (r *DBTxRun) Select(dest any, query string, args ...any) error {
	return r.DBoperator.Select(dest, query, args...)
}
