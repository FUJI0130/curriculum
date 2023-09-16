package errors

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int
}

func (e *CustomError) Error() string {
	return e.Message
}
