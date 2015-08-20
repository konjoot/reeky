package errors

type ValidationError struct{}

func (e *ValidationError) Error() string {
	return "ValidationError"
}
