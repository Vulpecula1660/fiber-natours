package middleware

import "github.com/gofiber/fiber/v2"

func ErrorHandler(err error) *fiber.Error {
	return fiber.NewError(50000, err.Error())
}
