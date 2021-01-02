package cerrors

import (
	"github.com/pkg/errors"
)

type custom struct {
	code Code
	err  error
}

func (e *custom) Error() string {
	return e.err.Error()
}

func Errorf(c Code, format string, args ...interface{}) error {
	return &custom{
		code: c,
		err:  errors.Errorf(format, args...),
	}
}

func GetCode(err error) Code {
	if err == nil {
		return OK
	}
	if e, ok := err.(*custom); ok {
		return e.code
	}
	return Unknown
}
