package base

type ErrorDetails struct {
	Message     string
	StatusCode  int
	Description string
}

type BaseError struct {
	details ErrorDetails
}

func NewBaseError(message string, statusCode int, description string) *BaseError {
	return &BaseError{
		details: ErrorDetails{
			Message:     message,
			StatusCode:  statusCode,
			Description: description,
		},
	}
}

func (be *BaseError) Error() string {
	return be.details.Message
}

func (be *BaseError) StatusCode() int {
	return be.details.StatusCode
}

func (be *BaseError) Details() ErrorDetails {
	return be.details
}
