package model

import (
	"fmt"
)

type ExtendedError struct {
	Cause []error
	text  string
}

func (e ExtendedError) Error() string {

	if e.text == "" {
		return fmt.Sprintf(e.Cause[0].Error())
	} else {
		return e.text
	}
}

func NewRequestError(text string, cause []error) *ExtendedError {
	e := ExtendedError{text: text, Cause: cause}
	return &e
}
