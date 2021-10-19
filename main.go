package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/v1/notes", getNotes)

	app.Listen(":8400")
}

func getNotes(c *fiber.Ctx) error {
	err := c.SendString("All good!")
	return err
}
