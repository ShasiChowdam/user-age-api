package main

import (
	"log"

	"github.com/ShasiChowdam/user-age-api/config"
	loggerpkg "github.com/ShasiChowdam/user-age-api/internal/logger"
	"github.com/ShasiChowdam/user-age-api/db/sqlc"
	"github.com/ShasiChowdam/user-age-api/internal/handler"
	"github.com/ShasiChowdam/user-age-api/internal/repository"
	"github.com/ShasiChowdam/user-age-api/internal/routes"
	"github.com/ShasiChowdam/user-age-api/internal/service"
	"github.com/ShasiChowdam/user-age-api/internal/middleware"
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

	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	appLogger.Info("Database connected successfully")

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	routes.SetupRoutes(app, userHandler)

	appLogger.Info(
		"Server started",
		zap.String("port", cfg.AppPort),
	)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}