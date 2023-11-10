package middleware

import (
	"log"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TransactionHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isModifyingMethod(c.Request.Method) {
			log.Printf("method is : %s\n", c.Request.Method)
			tx := handleTransaction(db, c)
			defer finalizeTransaction(tx, c)
			c.Set("Conn", tx)
			log.Printf("Transaction set in context: %v", tx != nil)

		} else {
			c.Set("Conn", db)
		}
		c.Next()
	}
}
func isModifyingMethod(method string) bool {
	log.Printf("method is : %s\n", method)
	switch method {
	case "POST", "PATCH", "PUT", "DELETE":
		return true
	default:
		return false
	}
}

func handleTransaction(db *sqlx.DB, c *gin.Context) *sqlx.Tx {
	log.Printf("handleTransaction")
	tx, err := db.Beginx()
	if err != nil {
		log.Printf("Failed to start transaction")
		wrappedErr := customerrors.WrapInternalServerError(err, "Failed to start transaction")
		c.Error(wrappedErr)
		return nil
	}
	return tx
}

func finalizeTransaction(tx *sqlx.Tx, c *gin.Context) {
	log.Printf("finalizeTransaction")
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	}

	if hasErrors(c) {
		log.Println("Errors detected, rolling back transaction.")
		tx.Rollback()
	} else if c.Writer.Status() >= 400 {
		log.Printf("HTTP Status %d detected, rolling back transaction.", c.Writer.Status())
		tx.Rollback()
	} else {
		log.Println("Committing transaction.")
		tx.Commit()
	}
}

func hasErrors(c *gin.Context) bool {
	ginErr, exists := c.Get("C_ERRORS")
	if !exists {
		log.Println("No errors detected in context.")
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
