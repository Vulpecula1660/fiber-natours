package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/Vulpecula1660/fiber-natours/api"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
	"github.com/Vulpecula1660/fiber-natours/enum"
)

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Return from handler
			customError := err.(*enum.CustomError)
			return ctx.Status(customError.HTTPStatus).JSON(protocol.Response{
				Code:    strconv.Itoa(customError.Code),
				Message: customError.Message,
				Result:  struct{}{},
			})
		},
	})

	app.Use(logger.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	api.SetupRoutes(app)

	app.Listen("127.0.0.1:3000")
}
