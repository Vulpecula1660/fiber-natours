package middleware

import (
	"fmt"

	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/redis"

	"github.com/gofiber/fiber/v2"
)

func Token(c *fiber.Ctx) error {
	token := c.Get("api-token")

	if token == "" {
		return enum.NonLogin
	}

	uuid, err := redis.Get(c.Context(), fmt.Sprintf("token:%s", token))
	if err != nil {
		return ErrorHandler(err)
	}

	if uuid == "" {
		return enum.NonLogin
	}

	c.Set("userID", uuid)

	return c.Next()
}
