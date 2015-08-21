package errors

func NewEmptyResourceError() *EmptyResourceError {
	return &EmptyResourceError{}
}

type EmptyResourceError struct{}

func (e *EmptyResourceError) Error() string {
	return "EmptyResourceError"
}
