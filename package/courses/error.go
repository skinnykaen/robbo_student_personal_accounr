package courses

import "errors"

var (
	ErrIncorrectInputParam = errors.New("user not found")
	ErrOnReq               = errors.New("error on request")
	ErrReadRespBody        = errors.New("error while reading the response bytes")
	ErrTknNotRefresh       = errors.New("token not refresh")
	ErrOnResp              = errors.New("error on response")
	ErrJsonMarshal         = errors.New("error on json Marshal")
	ErrCourseNotFound      = errors.New("course not found")
)
