package leakerrs

import (
	"time"
)

type Error struct {
	Msg  string
	Code int
	Time time.Time
}

func (e *Error) Error() string {
	return e.Msg
}

func new(msg string, code int) Error {
	return Error{
		Msg:  msg,
		Code: code,
		Time: time.Now(),
	}
}

func GetError(err error) Error {
	switch err.Error() {
	case DocumentNotFound:
		return new(err.Error(), 404)

	case DocumentExist:
		return new(err.Error(), 403)

	case InternalError:
		return new(err.Error(), 500)

	case CantConnectToDB:
		return new(err.Error(), 500)
	}

	return new(err.Error(), 500)
}

var (
	DocumentNotFound = "Docuement not found"
	DocumentExist    = "Document Exist"
	InternalError    = "Internal server err"
	CantConnectToDB  = "Cant connect to db"
)
