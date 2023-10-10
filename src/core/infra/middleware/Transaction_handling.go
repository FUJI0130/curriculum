package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TransactionHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var tx *sqlx.Tx
		var err error
		log.Printf("TransactionHandler: %s", c.Request.Method)

		// POST, PATCH, PUT, DELETE の場合トランザクションを開始
		switch c.Request.Method {
		case "POST", "PATCH", "PUT", "DELETE":
			tx, err = db.Beginx()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to start transaction"})
				return
			}

			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
					panic(r)
				} else if c.Writer.Status() >= 400 {
					tx.Rollback()
				} else {
					tx.Commit()
				}
			}()

			log.Printf("TransactionHandler before c.Request")
			// c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "tx", tx))
			c.Set("tx", tx) // トランザクションをgin.Contextに設定
		}

		c.Next()
	}
}
