package internal

type ErrParser struct {
	message string
}

func (e *ErrParser) Error() string {
	return e.message
}

type ErrParser2 struct {
	error
    limit int
	message string
}