package router

import (
	noteRoutes "github.com/galifornia/go-fiber-note-taking/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())

	noteRoutes.SetupNoteRoutes(&api)
}
