package main

import (
	"log"

	"github.com/ShasiChowdam/user-age-api/config"
	loggerpkg "github.com/ShasiChowdam/user-age-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	appLogger, err := loggerpkg.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer appLogger.Sync()

	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()

	appLogger.Info("Database connected successfully")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "User Age API Running",
		})
	})

	appLogger.Info(
		"Server started",
		zap.String("port", cfg.AppPort),
	)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}