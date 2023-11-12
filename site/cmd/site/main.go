package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	renderEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: renderEngine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "index")
	})

	app.Listen(":5555")
}
