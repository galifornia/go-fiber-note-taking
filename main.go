package main

import (
	"github.com/galifornia/go-fiber-note-taking/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Listen(":8400")
}
