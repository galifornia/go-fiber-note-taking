package noteRoutes

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router *fiber.Router) {
	note := (*router).Group("/note")

	// Create note
	note.Post("/", func(c *fiber.Ctx) error {
		var err error
		return err
	})

	// Delete note
	note.Delete("/:id", func(c *fiber.Ctx) error {
		var err error
		return err
	})

	// Get note
	note.Get("/:id", func(c *fiber.Ctx) error {
		var err error
		return err
	})

	// Get all notes
	note.Get("/", func(c *fiber.Ctx) error {
		var err error
		return err
	})

	// Update note
	note.Put("/:id", func(c *fiber.Ctx) error {
		var err error
		return err
	})
}
