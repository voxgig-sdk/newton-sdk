package core

type NewtonError struct {
	IsNewtonError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewNewtonError(code string, msg string, ctx *Context) *NewtonError {
	return &NewtonError{
		IsNewtonError: true,
		Sdk:              "Newton",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *NewtonError) Error() string {
	return e.Msg
}
