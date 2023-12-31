package noteHandler

import (
	"github.com/MikunsHub/fiberPlayground/database"
	"github.com/MikunsHub/fiberPlayground/internals/model"
	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
    db := database.DB
    var notes []model.Note

    // find all notes in the database
    db.Find(&notes)

    // If no note is present return an error
    if len(notes) == 0 {
        return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
    }

    // Else return notes
    return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
    db := database.DB
    note := new(model.Note)

    // Store the body in the note and return error if encountered
    err := c.BodyParser(note)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
    }
    // Add a uuid to the note
    // note.ID = uuid.New()
    // Create the Note and return error if encountered
    err = db.Create(&note).Error
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
    }

    // Return the created note
    return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}