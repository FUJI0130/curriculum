package customerrors

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
)

type BaseErr struct {
	Message       string
	StatusCodeVal int
	TraceVal      error
}

type BaseError interface {
	Error() string
	StatusCode() int
	Trace() error
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

	return wrappedError
}

// func (be *BaseErr) Error() string {
// 	stackTraceFilter := &support.StackTraceFilter{}
// 	traceString := fmt.Sprintf("%+v", be.TraceVal)
// 	var resultStackTrace = ""
// 	if config.GlobalConfig.DebugMode {
// 		resultStackTrace = stackTraceFilter.RemoveLinesFromKeywords(traceString)
// 	} else {
// 		resultStackTrace = traceString
// 	}
// 	lines := strings.SplitN(resultStackTrace, "\n", 2)
// 	if len(lines) > 1 {
// 		resultStackTrace = lines[1]
// 	}
// 	return fmt.Sprintf("%s ### \n%s", be.Message, resultStackTrace)
// }

func (be *BaseErr) Error() string {
	traceString := fmt.Sprintf("%+v", be.TraceVal)
	return fmt.Sprintf("%s ### \n%s", be.Message, traceString)
}

func SplitMessageAndTrace(errStr string) (string, string) {
	parts := strings.SplitN(errStr, " ### ", 2)
	if len(parts) < 2 {
		return errStr, ""
	}

	message := parts[0]

	stackTrace := parts[1]

	return message, stackTrace
}

func (be *BaseErr) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseErr) Trace() error {
	return be.TraceVal
}
