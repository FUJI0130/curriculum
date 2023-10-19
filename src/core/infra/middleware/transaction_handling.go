package middleware

import (
	"log"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
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

		switch c.Request.Method {
		case "POST", "PATCH", "PUT", "DELETE":
			tx, err = db.Beginx()
			if err != nil {
				wrappedErr := customerrors.WrapInternalServerError(err, "Failed to start transaction")
				c.Error(wrappedErr)
				return
			}

			defer func() {
				if r := recover(); r != nil {
					log.Printf("TransactionHandler: panic")
					tx.Rollback()
					log.Printf("TransactionHandler: rollback1")
					panic(r)
				} else {
					log.Printf("TransactionHandler: check error")
					// エラーがBaseErrorタイプであるかチェック
					if ginErr, exists := c.Get("C_ERRORS"); exists {
						if errArray, ok := ginErr.([]error); ok {
							for _, e := range errArray {
								if customErr, ok := e.(customerrors.BaseError); ok {
									log.Printf("TransactionHandler: error status code = %d", customErr.StatusCode())
									if customErr.StatusCode() >= 400 {
										log.Printf("TransactionHandler: custom error with status >= 400")
										tx.Rollback()
										log.Printf("TransactionHandler: rollback due to custom error")
										return
									}
								}

							}
						}
					}
					log.Printf("c.Writer.Status() = %d", c.Writer.Status())
					if c.Writer.Status() >= 400 {
						log.Printf("TransactionHandler: status >= 400")
						tx.Rollback()
						log.Printf("TransactionHandler: rollback2")
					} else {
						tx.Commit()
						log.Printf("TransactionHandler: commit")
					}
				}
			}()

			c.Set("tx", tx)
		}

		c.Next()
	}
}
