package cerrors

type Reason string

const (
	OK      Reason = "ok"
	Unknown Reason = "unknown"

	Duplicated       Reason = "duplicated"
	Unexpected       Reason = "un_expected"
	ValidationFailed Reason = "validation_Failed"
	DatabaseErr      Reason = "database_err"
)
