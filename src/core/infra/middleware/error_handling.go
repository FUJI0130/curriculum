package middleware

import (
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		handleError(c, c.Errors.Last().Err)
		return
	}
}

func handleError(c *gin.Context, err error) {
	if customErr, ok := err.(customerrors.BaseError); ok {
		message, _ := customerrors.SplitMessageAndTrace(customErr.Error())
		c.JSON(customErr.StatusCode(), gin.H{
			"message": message,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
