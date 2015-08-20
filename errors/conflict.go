package errors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "ConflictError"
}
