package helper

// ErrorFormat shows that it is impossible to reduce a string to a number
type ErrorFormat struct{}

func (e *ErrorFormat) Error() string {
	return "the number must contain only digits"
}
