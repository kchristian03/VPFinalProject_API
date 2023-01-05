package NoteController

import (
	"ZenZen_API/app/models"
	"ZenZen_API/framework/utils"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateNote struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// CreateNoteHandler
// This is an example of a handler
// Modify it to suit your needs
func CreateNoteHandler(c *fiber.Ctx) error {
	var note CreateNote
	_ = json.Unmarshal(c.Body(), &note)

	err := note.performValidation()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}

	newNote := models.Note{
		Title:   note.Title,
		Content: note.Content,
		UserID:  note.UserID,
	}

	createdNote, err := newNote.CreateNote()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Note " + createdNote.Title + " created successfully",
		Data:    createdNote,
	})

}

func (u CreateNote) performValidation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Title, validation.NotNil),
		validation.Field(&u.Content, validation.NotNil),
	)
}
