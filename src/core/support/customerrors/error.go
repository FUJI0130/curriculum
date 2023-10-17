package customerrors

import (
	"fmt"

	"github.com/FUJI0130/curriculum/src/core/config"
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
	// _, file, line, ok := runtime.Caller(2)
	// if !ok {
	// 	message = fmt.Sprintf("Failed to get runtime caller information: %s", message)
	// } else {
	// 	message = fmt.Sprintf("at [File: %s Line: %d] %s", file, line, message)
	// }

	wrappedError := &BaseErr{
		Message:       message,
		StatusCodeVal: b.StatusCodeVal,
		TraceVal:      errors.Wrap(err, message),
	}

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

// func (be *BaseErr) Error() string {
// 	return be.Message
// }

// スタックトレース出すのに必要
//
//	func (be *BaseErr) Error() string {
//		return fmt.Sprintf("%s: %+v", be.Message, be.TraceVal)
//	}
func (be *BaseErr) Error() string {
	return fmt.Sprintf("%s ### %v", be.Message, be.TraceVal)
}

func (be *BaseErr) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseErr) Trace() error {
	return be.TraceVal
}
