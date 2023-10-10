package customerrors

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

func (be *BaseErr) Error() string {
	return be.Message
}

func (be *BaseErr) StatusCode() int {
	return be.StatusCodeVal
}

func (be *BaseErr) Trace() error {
	return be.TraceVal
}
