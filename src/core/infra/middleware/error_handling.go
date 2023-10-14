package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

// func ErrorHandler(c *gin.Context) {
// 	defer func() {
// 		if rec := recover(); rec != nil {
// 			var err error
// 			if casted, ok := rec.(error); ok {
// 				err = casted
// 			} else {
// 				err = errors.New("unknown panic")
// 			}
// 			log.Printf("recovered from panic: %+v", err)
// 			log.Printf("Recovered error type: %T, value: %+v", err, err)
// 			switch e := err.(type) {
// 			case customerrors.BaseError:
// 				log.Printf("ERROR: %+v", e.Trace())
// 				c.JSON(e.StatusCode(), gin.H{
// 					"message": fmt.Sprintf("%d: An error occurred", e.StatusCode()), // Simple error message to the client
// 				})
// 			default:
// 				log.Printf("FATAL: %+v", e)
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"message": "Fatal",
// 				})
// 			}
//				c.Abort()
//			}
//		}()
//		c.Next()
//	}

func ErrorHandler(c *gin.Context) {
	defer func() {
		if rec := recover(); rec != nil {
			handleError(c, convertToError(rec))
		}
	}()

	c.Next() // Execute the rest of middlewares & route handler, errors are collected into c.Errors.

	// Handle error if c.Errors is non-nil (non-empty)
	for _, err := range c.Errors {
		handleError(c, err.Err)
	}
}

func handleError(c *gin.Context, err error) {
	log.Printf("Error: %+v", err)

	// Here, wrapping the original error with stack trace info.
	if isDebugMode() {
		err = errors.WithStack(err)
		log.Printf("Debug Mode - Stack Trace: %+v", err)
	}

	if e, ok := err.(customerrors.BaseError); ok {
		// err is of type customerrors.BaseError
		c.JSON(e.StatusCode(), gin.H{
			"message": e.Error(),
		})
		return
	}
	// Default error handling
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
