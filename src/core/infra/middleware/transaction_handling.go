package middleware

import (
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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
		handleTransactionError(c, err)
		return nil
	}
	c.Set("tx", tx)
	return tx
}
func handleTransactionError(c *gin.Context, err error) {
	wrappedErr := customerrors.WrapInternalServerError(err, "Failed to start transaction")
	c.Error(wrappedErr)
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
