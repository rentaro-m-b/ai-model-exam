package request

type ValidationError int

const (
	ValidationErrRequestFieldMissing ValidationError = iota
	ValidationErrRequestFieldEmpty
)
