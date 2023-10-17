package middleware

import (
	"net/http"
	"strings"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // Execute the rest of middlewares & route handler, errors are collected into c.Errors.

	for _, err := range c.Errors {
		handleError(c, err.Err)
	}
}
func handleError(c *gin.Context, err error) {
	// custom errorの場合、そのエラーメッセージを使用してレスポンスを返す
	if customErr, ok := err.(customerrors.BaseError); ok {
		// フィルタリングしたスタックトレースの最初の行を取得
		firstLineOfTrace := strings.SplitN(customErr.Error(), "\n", 2)[0]

		// 最初の行からエラーメッセージとスタックトレースを分割
		message, _ := splitErrorAndStackTrace(firstLineOfTrace)
		c.JSON(customErr.StatusCode(), gin.H{
			// "message": customErr.Error(),
			"message": message,
		})
		return
	}

	// それ以外の場合、内部サーバーエラーとしてレスポンスを返す
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}

// func handleError(c *gin.Context, err error) {
// 	stackTraceFilter := &StackTraceFilter{}
// 	// 最初の5行を取得
// 	fullStackTrace := stackTraceFilter.ExtractNLinesFromStart(err, 5)
// 	// 最初の行を取得
// 	firstLineOfTrace := strings.SplitN(fullStackTrace, "\n", 2)[0]
// 	// 最初の行からエラーメッセージとスタックトレースを分割
// 	message, extractedTrace := splitErrorAndStackTrace(firstLineOfTrace)
// 	log.Printf("Filtered Stack Trace (length: %d): %v", len(fullStackTrace), extractedTrace)
// 	// custom errorの場合、そのエラーメッセージを使用してレスポンスを返す
// 	if customErr, ok := err.(customerrors.BaseError); ok {
// 		// log.Println(customErr.Trace())
// 		c.JSON(customErr.StatusCode(), gin.H{
// 			"message": message,
// 		})
// 		return
// 	}
// 	// それ以外の場合、内部サーバーエラーとしてレスポンスを返す
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "Internal Server Error",
// 	})
// }

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
