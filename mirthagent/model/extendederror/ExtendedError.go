package extendederror

import (
	"fmt"
)

type ExtendedError interface {
	Error() string
	Cause() []error
}

type extendedError struct {
	cause []error
	text  string
}

func (e extendedError) Error() string {
	if e.text == "" {
		return fmt.Sprintf(e.cause[0].Error())
	} else {
		return e.text
	}
}

func (e extendedError) Cause() []error {
	return e.cause
}

func New(text string, cause []error) *extendedError {
	e := extendedError{text: text, cause: cause}
	return &e
}
