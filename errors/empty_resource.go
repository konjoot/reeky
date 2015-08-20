package errors

type EmptyResourceError struct{}

func (e *EmptyResourceError) Error() string {
	return "EmptyResourceError"
}
