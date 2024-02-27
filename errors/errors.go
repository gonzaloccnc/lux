package errors

type UserAborted struct{}

func (e *UserAborted) Error() string {
	return "user quit prompt"
}
