package base

type BaseError struct {
	Message       string
	StatusCodeVal int
	TraceVal      error
}

type BaseErrorHandler interface {
	StatusCode() int
	Trace() error
	Error() string
}

func NewBaseError(message string, statusCode int, trace error) *BaseError {
	return &BaseError{
		Message:       message,
		StatusCodeVal: statusCode,
		TraceVal:      trace,
	}
}

func (be *BaseError) Error() string {
	return be.Message
}

func (be *BaseError) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseError) Trace() error {
	return be.TraceVal
}
