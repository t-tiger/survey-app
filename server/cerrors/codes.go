package cerrors

type Reason int

const (
	OK Reason = iota
	Unknown
	Duplicated
	InvalidInput
	Unauthorized
	Unexpected
	DatabaseErr
)
