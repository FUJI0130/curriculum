package middleware

import (
	"log"
	"net/http"
	"strings"

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
	// エラーメッセージの文字数を取得
	// errorLength := len(fmt.Sprintf("%+v", err))

	// log.Printf("Error (length: %d): %+v", errorLength, err)

	// エラーログの出力部分でStackTraceをフィルタリングする
	stackTraceFilter := &StackTraceFilter{}
	log.Printf("Before Filtering: %+v", err)
	filteredStackTrace := stackTraceFilter.FilterStackTrace(err)
	log.Printf("After Filtering: %v", filteredStackTrace)

	// フィルタリング後のスタックトレースの文字数を取得
	filteredLength := len(filteredStackTrace)

	// log.Printf("Filtered Stack Trace (length: %d): %v", filteredLength, filteredStackTrace)
	// log.Printf("Filtered Stack Trace (length: %d): ", filteredLength)

	// スタックトレースから1行目だけ取得
	firstLine := strings.Split(filteredStackTrace, "\n")[0]

	// message := customerrors.ExtractMessageFromError(firstLine)
	message, stackTrace := splitErrorAndStackTrace(firstLine)

	log.Printf("Filtered Stack Trace (length: %d): %v", filteredLength, stackTrace)
	// log.Printf("Non Filtered Stack Trace (length: %d): %v", filteredLength, err)

	if e, ok := err.(customerrors.BaseError); ok {
		c.JSON(e.StatusCode(), gin.H{
			// "message": e.Error(),
			// "message": firstLine,
			"message": message,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
func splitErrorAndStackTrace(errStr string) (string, string) {
	parts := strings.SplitN(errStr, " ### ", 2)
	if len(parts) < 2 {
		return errStr, ""
	}

	message := parts[0]

	// スタックトレースの全体を取得するため、" ### " の後のすべての文字列を取得します。
	stackTrace := parts[1]

	return message, stackTrace
}
