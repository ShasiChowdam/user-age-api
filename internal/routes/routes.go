package routes

import (
	"github.com/ShasiChowdam/user-age-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.ListUsers)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Get("/users/:id", userHandler.GetUserByID)
}