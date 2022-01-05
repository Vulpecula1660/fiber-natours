package middleware

import (
	"net/http"

	"github.com/Vulpecula1660/fiber-natours/enum"
)

func ErrorHandler(err error) *enum.CustomError {
	// default error code
	errCode := 99999

	if e, ok := err.(*enum.CustomError); ok {
		return e
	}

	return &enum.CustomError{
		HTTPStatus: http.StatusInternalServerError,
		Code:       errCode,
		Message:    err.Error(),
	}
}
