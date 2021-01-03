package cerrors

type Reason int

const (
	OK Reason = iota
	Unknown
	DatabaseErr
	Duplicated
	InvalidInput
	Forbidden
	NotFound
	Unauthorized
	Unexpected
)
