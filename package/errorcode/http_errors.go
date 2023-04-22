package errorcode

import (
	"github.com/kondroid00/sample-server-2022/package/errors"
	e "github.com/pkg/errors"
)

type HttpError struct {
	httpStatusCode int
	errorCode      ErrorCode
	err            errors.Error
}

func NewHttpError(code int, errorCode ErrorCode, err errors.Error) errors.Error {
	return &HttpError{
		httpStatusCode: code,
		errorCode:      errorCode,
		err:            err,
	}
}

func (e *HttpError) HttpStatusCode() int {
	return e.httpStatusCode
}

func (e *HttpError) ErrorCode() ErrorCode {
	return e.errorCode
}

func (e *HttpError) Error() string {
	return e.err.Error()
}

func (e *HttpError) StackTrace() e.StackTrace {
	return e.err.StackTrace()
}

func (e *HttpError) Unwrap() error {
	return e.err.Unwrap()
}
