package NoteController

import (
	"ZenZen_API/app"
	"ZenZen_API/app/models"
	"ZenZen_API/framework/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type EditNote struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// EditNoteHandler
// This is an example of a handler
// Modify it to suit your needs
func EditNoteHandler(c *fiber.Ctx) error {

	var editNote EditNote
	_ = json.Unmarshal(c.Body(), &editNote)

	var note models.Note
	app.DB.First(&note, c.Params("id"))

	note.Title = editNote.Title
	note.Content = editNote.Content

	app.DB.Save(&note)

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Note " + note.Title + " edited successfully",
		Data:    note,
	})
}
