package customerrors

import (
	"fmt"
	"strings"

	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/FUJI0130/curriculum/src/core/support"
	"github.com/cockroachdb/errors"
)

type BaseError interface {
	Error() string
	StatusCode() int
	Trace() error
}

type BaseErr struct {
	Message       string
	StatusCodeVal int
	TraceVal      error
}

func NewBaseError(message string, statusCode int, trace error) *BaseErr {
	return &BaseErr{
		Message:       message,
		StatusCodeVal: statusCode,
		TraceVal:      trace,
	}
}

func (b *BaseErr) WrapWithLocation(err error, message string) *BaseErr {

	wrappedError := &BaseErr{
		Message:       message,
		StatusCodeVal: b.StatusCodeVal,
		TraceVal:      errors.Wrap(err, message),
	}

	if config.GlobalConfig.DebugMode {
		// fmt.Println(wrappedError.TraceVal)
	}
	return wrappedError
}

//	func (be *BaseErr) Error() string {
//		return fmt.Sprintf("%s ### %+v", be.Message, be.TraceVal)
//	}

func (be *BaseErr) Error() string {
	stackTraceFilter := &support.StackTraceFilter{}

	traceString := fmt.Sprintf("%+v", be.TraceVal)
	// log.Println("traceString is : ", traceString)

	// be.TraceValを直接フィルタリングします。
	// filteredTrace := stackTraceFilter.ExtractNLinesFromStart(traceString, 5)
	filteredTrace := stackTraceFilter.RemoveLinesFromKeywords(traceString)

	// フィルタリングしたスタックトレースの最初の行を取得
	firstLineOfTrace := strings.SplitN(filteredTrace, "\n", 2)[0]

	// 最初の行からエラーメッセージとスタックトレースを分割
	message, _ := splitErrorAndStackTrace(firstLineOfTrace)

	// このmessageを使用してエラーメッセージを構築
	// return fmt.Sprintf("%s ### %s", be.Message, message)
	return fmt.Sprintf("%s ### %s", message, filteredTrace)
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

func (be *BaseErr) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseErr) Trace() error {
	return be.TraceVal
}
