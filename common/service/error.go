package service

import "fmt"

type error interface {
	Error() string
}

type BadRequest struct {
	ErrMessage string
}

type InternalServer struct {
	ErrMessage string
}

type NotFound struct {
	ErrMessage string
}

type UnAuthorized struct {
	ErrMessage string
}

func (e *BadRequest) Error() string {
	return fmt.Sprintf("%v: BadRequestError", e.ErrMessage)
}

func (e *InternalServer) Error() string {
	return fmt.Sprintf("%v: InternalServerError", e.ErrMessage)
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%v: NotFoundError", e.ErrMessage)
}

func (e *UnAuthorized) Error() string {
	return fmt.Sprintf("%v: UnauthorizedError", e.ErrMessage)
}
