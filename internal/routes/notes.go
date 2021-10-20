package noteRoutes

import (
	noteHandler "github.com/galifornia/go-fiber-note-taking/internal/handlers/notes"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router *fiber.Router) {
	note := (*router).Group("/note")

	// Create note
	note.Post("/", noteHandler.CreateNote)

	// Delete note
	note.Delete("/:id", noteHandler.DeleteNote)

	// Get note
	note.Get("/:id", noteHandler.GetNote)

	// Get all notes
	note.Get("/", noteHandler.GetNotes)

	// Update note
	note.Put("/:id", noteHandler.UpdateNote)
}
