package cerrors

import (
	"github.com/pkg/errors"
)

type custom struct {
	reason Reason
	err    error
}

func (e *custom) Error() string {
	return e.err.Error()
}

func Errorf(reason Reason, format string, args ...interface{}) error {
	return &custom{
		reason: reason,
		err:    errors.Errorf(format, args...),
	}
}

func GetReason(err error) Reason {
	if err == nil {
		return OK
	}
	if e, ok := err.(*custom); ok {
		return e.reason
	}
	return Unknown
}
