package config

import "github.com/gofiber/fiber/v2"

func Views(viewName string, data fiber.Map) fiber.Handler {
	viewName = "views/" + viewName
	return func(c *fiber.Ctx) error {
		return c.Render(viewName, data)
	}
}

func ViewsWithLayout(viewName string, data fiber.Map, layoutName string) fiber.Handler {
	viewName = "views/pages/" + viewName
	layoutName = "views/layouts/" + viewName
	return func(c *fiber.Ctx) error {
		return c.Render(viewName, data, layoutName)
	}
}
