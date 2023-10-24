package rdbimpl

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SqlxTransaction struct {
	Tx *sqlx.Tx
}
type SQLResult struct {
	res sql.Result
}
type Transaction interface {
	Exec(query string, args ...any) (Result, error)
	Commit() error
	Rollback() error
}
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

func (t *SqlxTransaction) Commit() error {
	return t.Tx.Commit()
}

func (t *SqlxTransaction) Rollback() error {
	return t.Tx.Rollback()
}
func (r *SQLResult) LastInsertId() (int64, error) {
	return r.res.LastInsertId()
}

func (r *SQLResult) RowsAffected() (int64, error) {
	return r.res.RowsAffected()
}

func (t *SqlxTransaction) Exec(query string, args ...any) (Result, error) {
	result, err := t.Tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &SQLResult{res: result}, nil
}
