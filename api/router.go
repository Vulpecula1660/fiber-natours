package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Vulpecula1660/fiber-natours/api/controller/front/member"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// 前台 API - 未登入
	v1FrontNonToken := app.Group("/v1/front")
	{
		v1FrontNonToken.Post("/member/register", member.Register) // 註冊
		v1FrontNonToken.Post("/member/login", member.Login)       // 登入
	}

	// routes
	// api.Get("/", handler.GetAllProducts)
	// api.Get("/:id", handler.GetSingleProduct)
	// api.Post("/", handler.CreateProduct)
	// api.Delete("/:id", handler.DeleteProduct)
}
