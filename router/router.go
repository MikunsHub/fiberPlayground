package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "github.com/MikunsHub/fiberPlayground/internals/routes/note"
)

func SetupRoutes(app *fiber.App) {
	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
    noteRoutes.SetupNoteRoutes(api)
}