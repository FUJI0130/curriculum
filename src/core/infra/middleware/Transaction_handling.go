package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func TransactionHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		existingTx, _ := c.Get("tx")
		if existingTx != nil {
			c.Next()
			return
		}

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
					log.Printf("TransactionHandler: panic")
					tx.Rollback()
					log.Printf("TransactionHandler: rollback1")
					panic(r)
				} else if c.Writer.Status() >= 400 {
					log.Printf("TransactionHandler: status >= 400")
					tx.Rollback()
					log.Printf("TransactionHandler: rollback2")
				} else {
					tx.Commit()
					log.Printf("TransactionHandler: commit")
				}
			}()

			log.Printf("TransactionHandler before c.Request")
			// c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "tx", tx))
			c.Set("tx", tx) // トランザクションをgin.Contextに設定
		}

		c.Next()
	}
}
