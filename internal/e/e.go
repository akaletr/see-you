package e

import "fmt"

type e struct {
	Err  error
	code errorCode
}

func (e *e) Error() string {
	return fmt.Sprintf("Error (%s): %s", e.code, e.Err())
}

func New(err error, code errorCode) error {
	return &e{
		Err:  err,
		code: code,
	}
}
