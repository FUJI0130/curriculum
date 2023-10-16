package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/config"
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

	if isDebugMode() {
		err = errors.WithStack(err)
		log.Printf("Debug Mode - Stack Trace: %+v", err)
	}
	// Assuming here that you'd handle panics in a generic way, e.g., returning a 500.
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}

func convertToError(r interface{}) error {
	if err, ok := r.(error); ok {
		return err
	}
	return errors.New(fmt.Sprint(r))
}

func isDebugMode() bool {
	return config.GlobalConfig.DebugMode
}
