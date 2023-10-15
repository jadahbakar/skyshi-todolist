package errorlib

import "fmt"

type ErrorCode uint

type Error struct {
	origin error
	code   ErrorCode
	msg    string
}

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
	ErrorCodePreconditionFailed
	ErrorCodeValidationError
	ErrorCodeUnauthorized
	ErrorCodeInternal
	ErrorCodeBadRequest
)

func WrapErr(origin error, code ErrorCode, message string, a ...interface{}) error {
	return &Error{
		origin: origin,
		code:   code,
		msg:    fmt.Sprintf(message, a...),
	}
}

// NewErrorf instantiates a new error.
func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapErr(nil, code, format, a...)
}

func (e *Error) Error() string {
	if e.origin != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.origin)
	}
	return e.msg
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	return e.origin
}

// Code returns the code representing this error.
func (e *Error) Code() ErrorCode {
	return e.code
}
