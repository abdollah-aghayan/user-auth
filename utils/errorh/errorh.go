package errorh

import "net/http"

//Errorh struct
type Errorh struct {
	error
	Message string
	Code    int
}

//InternalError internal error
func InternalError(msg string) *Errorh {
	return newError(msg, http.StatusInternalServerError)
}

//BadRequestError bad request error
func BadRequestError(msg string) *Errorh {
	return newError(msg, http.StatusBadRequest)
}

// NotFoundError not found error
func NotFoundError(msg string) *Errorh {
	return newError(msg, http.StatusNotFound)
}

// NotAuthorizedError not authorized error
func NotAuthorizedError(msg string) *Errorh {
	return newError(msg, http.StatusUnauthorized)
}

func newError(msg string, code int) *Errorh {
	return &Errorh{
		Code:    code,
		Message: msg,
	}
}

func (e *Errorh) Error() string {
	return e.Message
}
