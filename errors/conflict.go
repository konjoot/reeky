package errors

func NewConflictError() *ConflictError {
	return &ConflictError{}
}

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "ConflictError"
}
