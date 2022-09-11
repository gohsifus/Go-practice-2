// Package errs Для определения пользовательских ошибок
package errs

type errorType uint

const (
	UnHandledErr     = errorType(500)
	IncorrectDataErr = errorType(400)
	BusinessLogicErr = errorType(503)
)

type CustomError struct {
	status        errorType
	originalError error
}

func (c CustomError) Error() string {
	return c.originalError.Error()
}

func New(err error, status errorType) CustomError {
	return CustomError{
		status:        status,
		originalError: err,
	}
}

func (c CustomError) Status() int {
	return int(c.status)
}

func Wrap(err error) CustomError {
	if err, ok := err.(CustomError); ok {
		return err
	} else {
		return New(err, UnHandledErr)
	}
}
