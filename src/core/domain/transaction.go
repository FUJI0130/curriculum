package domain

type Transaction interface {
	Exec(query string, args ...any) (Result, error)
	Commit() error
	Rollback() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
