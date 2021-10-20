package noteHandler

import (
	"fmt"

	"github.com/galifornia/go-fiber-note-taking/database"
	"github.com/galifornia/go-fiber-note-taking/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all GetNotes
	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes found!", "data": notes})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body from request in created note
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot parse body data.", "data": err})

	}

	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong creating the note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note was created succesfully", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var note model.Note

	db.Find(&note, "id= ?", id)

	if note.ID == uuid.Nil {
		message := fmt.Sprintf("Note with id %s was not found", id)
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": message, "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note was found!", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	// Find note
	var note model.Note
	db.Find(&note, "id = ?", id)

	// Note not found
	if note.ID == uuid.Nil {
		message := fmt.Sprintf("Note with id %s was not found", id)
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": message, "data": nil})
	}

	// Delete note
	err := db.Delete(&note, "id = ?", id).Error
	if err != nil {
		message := fmt.Sprintf("Something went wrong deleting the note with id: %s", id)
		c.Status(404).JSON(fiber.Map{"status": "error", "message": message, "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note was deleted"})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB
	id := c.Params("id")

	var note model.Note
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		message := fmt.Sprintf("Note with id %s was not found", id)
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": message})
	}

	// Temporarely hold the value of the note
	var aux updateNote

	err := c.BodyParser(&aux)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while parsing input from request", "data": err})
	}

	note.Title = aux.Title
	note.Subtitle = aux.SubTitle
	note.Text = aux.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Note was updated", "data": note})
}
