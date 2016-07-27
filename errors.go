package bowling

type emptyArrayError struct {
	message string
}

func (e *emptyArrayError) Error() string {
	return e.message
}