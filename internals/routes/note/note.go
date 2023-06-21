package noteRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/MikunsHub/fiberPlayground/internals/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", noteHandler.CreateNotes)
	// Read all Notes
	note.Get("/", noteHandler.GetNotes)
	// Read one Note
	// note.Get("/:noteId", func(c *fiber.Ctx) error {})
	// // Update one Note
	// note.Put("/:noteId", func(c *fiber.Ctx) error {})
	// // Delete one Note
	// note.Delete("/:noteId", func(c *fiber.Ctx) error {})
}


