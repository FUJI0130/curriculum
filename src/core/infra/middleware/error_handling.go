package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
	}()
	c.Next()
}
