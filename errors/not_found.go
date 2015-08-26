package errors

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "NotFoundError"
}
