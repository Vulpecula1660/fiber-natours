package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(err error) *fiber.Error {
	// default error code
	errCode := 99999

	if e, ok := err.(*fiber.Error); ok {
		if e.Code != 0 {
			errCode = e.Code
		}

		return fiber.NewError(errCode, e.Message)
	}

	return fiber.NewError(errCode, err.Error())
}
