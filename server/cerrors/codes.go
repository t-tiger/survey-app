package cerrors

type Code string

const (
	OK      Code = "ok"
	Unknown Code = "unknown"

	Duplicated       Code = "duplicated"
	Unexpected       Code = "un_expected"
	ValidationFailed Code = "validation_Failed"
	DatabaseErr      Code = "database_err"
)
