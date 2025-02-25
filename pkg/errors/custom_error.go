package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type CustomError struct {
	code    string
	status  Status
	message string
}

const repoCode = "77777"

func NewCustomError(code string, status Status, message string) CustomError {
	return CustomError{
		code:    fmt.Sprintf("%s%s", repoCode, code),
		status:  status,
		message: message,
	}
}

func CauseCustomError(err error) CustomError {
	type causer interface {
		Cause() error
	}

	var result CustomError

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		customError, ok := err.(CustomError)
		if ok {
			result = customError
		}

		err = cause.Cause()
	}

	return result
}

func (e CustomError) Error() string {
	return e.message
}

func (e CustomError) Code() string {
	return e.code
}

func (e CustomError) Status() Status {
	return e.status
}

func (e CustomError) Wrap(err error) error {
	return errors.WithStack(e)
}
