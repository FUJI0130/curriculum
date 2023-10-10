package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/support/base"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			var err error
			switch t := rec.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("unknown panic")
			}
			log.Printf("recovered from panic: %+v", err) // %+v でstack traceもログに出力

			switch e := err.(type) {
			case base.BaseErrorHandler:
				log.Printf("ERROR: %+v", e.Trace())
				c.JSON(e.StatusCode(), gin.H{
					"message": fmt.Sprintf("%d: %s", e.StatusCode(), e.Error()),
				})
			default:
				log.Printf("FATAL: %+v", e)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Fatal",
				})
			}

			c.Abort()
		}
	}()
	c.Next()
}
