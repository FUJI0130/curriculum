package middleware

import (
	"database/sql"

	"github.com/FUJI0130/curriculum/src/core/domain"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SQLResult struct {
	res sql.Result
}

func (r *SQLResult) LastInsertId() (int64, error) {
	return r.res.LastInsertId()
}

func (r *SQLResult) RowsAffected() (int64, error) {
	return r.res.RowsAffected()
}

type SqlxTransaction struct {
	tx *sqlx.Tx
}

func (s *SqlxTransaction) Exec(query string, args ...any) (domain.Result, error) {
	result, err := s.tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &SQLResult{res: result}, nil
}

func (s *SqlxTransaction) Commit() error {
	return s.tx.Commit()
}

func (s *SqlxTransaction) Rollback() error {
	return s.tx.Rollback()
}

func TransactionHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isModifyingMethod(c.Request.Method) {
			tx := handleTransaction(db, c)
			defer finalizeTransaction(tx, c)
		}
		c.Next()
	}
}
func isModifyingMethod(method string) bool {
	switch method {
	case "POST", "PATCH", "PUT", "DELETE":
		return true
	default:
		return false
	}
}

func handleTransaction(db *sqlx.DB, c *gin.Context) *sqlx.Tx {
	tx, err := db.Beginx()
	if err != nil {
		wrappedErr := customerrors.WrapInternalServerError(err, "Failed to start transaction")
		c.Error(wrappedErr)
		return nil
	}

	transaction := &SqlxTransaction{tx: tx}

	c.Set("transaction", transaction)
	return tx
}

func finalizeTransaction(tx *sqlx.Tx, c *gin.Context) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	}

	if hasErrors(c) || c.Writer.Status() >= 400 {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func hasErrors(c *gin.Context) bool {
	ginErr, exists := c.Get("C_ERRORS")
	if !exists {
		return false
	}

	errArray, ok := ginErr.([]error)
	if !ok {
		return false
	}

	for _, e := range errArray {
		if customErr, ok := e.(customerrors.BaseError); ok && customErr.StatusCode() >= 400 {
			return true
		}
	}
	return false
}
