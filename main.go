package main

import (
	"github.com/galifornia/go-fiber-note-taking/database"
	"github.com/galifornia/go-fiber-note-taking/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":8400")
}
