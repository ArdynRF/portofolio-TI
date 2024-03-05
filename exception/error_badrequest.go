package exception

type BadRequestError struct {
	Error string
}

func NewBadRequesError(error string) BadRequestError {
	return BadRequestError{Error: error}
}
