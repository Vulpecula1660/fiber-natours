package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(err *fiber.Error) *fiber.Error {
	// default error code
	var errCode = 99999

	if err.Code != 0 {
		errCode = err.Code
	}

	return fiber.NewError(errCode, err.Message)
}
