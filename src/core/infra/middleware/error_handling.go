package middleware

import (
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // Execute the rest of middlewares & route handler, errors are collected into c.Errors.

	// Handle error if c.Errors is non-nil (non-empty)
	for _, err := range c.Errors {
		handleError(c, err.Err)
	}
}

func handleError(c *gin.Context, err error) {
	log.Printf("Error: %+v", err)
	if e, ok := err.(customerrors.BaseError); ok {
		c.JSON(e.StatusCode(), gin.H{
			"message": e.Error(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
