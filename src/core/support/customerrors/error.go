package customerrors

import (
	"fmt"
	"runtime"

	"github.com/FUJI0130/curriculum/src/core/config"
	"github.com/cockroachdb/errors"
)

type BaseErr struct {
	Message       string
	StatusCodeVal int
	TraceVal      error
}

type BaseError interface {
	StatusCode() int
	Trace() error
	Error() string
}

func NewBaseError(message string, statusCode int, trace error) *BaseErr {
	return &BaseErr{
		Message:       message,
		StatusCodeVal: statusCode,
		TraceVal:      trace,
	}
}

func (b *BaseErr) WrapWithLocation(err error, message string) *BaseErr {
	_, file, line, ok := runtime.Caller(2) // Caller(2) で呼び出し元の情報を取得
	if !ok {
		message = fmt.Sprintf("Failed to get runtime caller information: %s", message)
	} else {
		message = fmt.Sprintf("at [File: %s Line: %d] %s", file, line, message)
	}

	wrappedError := &BaseErr{
		Message:       message,
		StatusCodeVal: b.StatusCodeVal,
		TraceVal:      errors.Wrap(err, message),
	}

	// DebugModeがONの場合、スタックトレースを表示
	if config.GlobalConfig.DebugMode {
		fmt.Println(wrappedError.TraceVal)
	}
	return wrappedError
}

func (b *BaseErr) LogStackTrace() {
	if config.GlobalConfig.DebugMode {
		fmt.Println(b.TraceVal)
	}
}
func (be *BaseErr) Error() string {
	return be.Message
}

func (be *BaseErr) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseErr) Trace() error {
	return be.TraceVal
}
