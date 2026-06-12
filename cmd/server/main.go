package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "User Age API Running",
		})
	})

	log.Println("Server started on port 8080")

	log.Fatal(app.Listen(":8080"))
}