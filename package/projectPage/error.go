package projectPage

import "errors"

var (
	ErrPageNotFound        = errors.New("page not found")
	ErrInternalServerLevel = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrBadRequestBody      = errors.New("bad request body")
)
