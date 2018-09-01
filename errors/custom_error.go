package errors

import (
	"fmt"
)

type ErrorWithCode struct{
	code int
	message string
}

func New(text string) error {
	return &ErrorWithCode{
		message: text,
	}
}

func NewWithCode(paramCode int, paramText string) error {
	return &ErrorWithCode{
		code: paramCode,
		message: paramText,
	}
}

func (e *ErrorWithCode) Error() string {
	return fmt.Sprintf("(%d) %s",e.code, e.message)
}

func (e *ErrorWithCode) Code() int {
	return e.code
}