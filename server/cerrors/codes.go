package cerrors

type Reason string

const (
	OK      Reason = "ok"
	Unknown Reason = "unknown"

	Duplicated   Reason = "duplicated"
	InvalidInput Reason = "invalid_input"
	Unexpected   Reason = "un_expected"
	DatabaseErr  Reason = "database_err"
)
