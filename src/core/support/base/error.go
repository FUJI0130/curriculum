package base

type BaseError struct {
	Message    string
	StatusCode int
	Trace      error
}

func NewBaseError(message string, statusCode int, trace error) *BaseError {
	return &BaseError{
		Message:    message,
		StatusCode: statusCode,
		Trace:      trace,
	}
}

func (be *BaseError) Error() string {
	return be.Message
}
