package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

func PanicRecovery(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			handlePanic(c, convertToError(rec))
		}
	}()

	c.Next()
}

func handlePanic(c *gin.Context, err error) {
	log.Printf("Recovered from panic: %+v", err)

	err = errors.WithStack(err)
	log.Printf("Stack Trace: %+v", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "handlePanic Internal Server Error",
	})
}

func convertToError(r interface{}) error {
	if err, ok := r.(error); ok {
		return err
	}
	return errors.New(fmt.Sprint(r))
}
