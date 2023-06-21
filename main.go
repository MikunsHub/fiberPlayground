package main

import (
	"github.com/MikunsHub/fiberPlayground/database"
	"github.com/MikunsHub/fiberPlayground/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
    database.ConnectDB()

	// Setup the router
    router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
