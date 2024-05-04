package config

import "github.com/gofiber/fiber/v2"

func Views(viewName string, data fiber.Map) fiber.Handler {
	viewName = "views/" + viewName
	return func(c *fiber.Ctx) error {
		return c.Render(viewName, data)
	}
}
