package constants

type HandleResult int

const (
	NoHandlerResult HandleResult = iota + 1
	PassHandlerResult
	RefuseHandlerResult
	CancelHandlerResult
)