package routes

import "github.com/gofiber/fiber/v2"

func AppRoute() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	app.Listen(":8080")

}
