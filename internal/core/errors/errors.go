package errors

import "fmt"

type Kind string

const (
	KindValidation    Kind = "validation"
	KindConfiguration Kind = "configuration"
	KindNotFound      Kind = "not_found"
	KindConflict      Kind = "conflict"
	KindInternal      Kind = "internal"
)

type Error struct {
	Kind    Kind
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.Err == nil {
		return e.Message
	}

	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}

func New(kind Kind, message string) error {
	return &Error{Kind: kind, Message: message}
}

func Wrap(kind Kind, message string, err error) error {
	return &Error{Kind: kind, Message: message, Err: err}
}

func ExitCode(err error) int {
	if err == nil {
		return 0
	}

	if typed, ok := err.(*Error); ok {
		switch typed.Kind {
		case KindValidation:
			return 2
		case KindConfiguration:
			return 3
		case KindNotFound:
			return 4
		case KindConflict:
			return 5
		default:
			return 1
		}
	}

	return 1
}
