package main

import (
	"log"

	"github.com/ShasiChowdam/user-age-api/config"
	"github.com/ShasiChowdam/user-age-api/db/sqlc"
	"github.com/ShasiChowdam/user-age-api/internal/handler"
	loggerpkg "github.com/ShasiChowdam/user-age-api/internal/logger"
	"github.com/ShasiChowdam/user-age-api/internal/middleware"
	"github.com/ShasiChowdam/user-age-api/internal/repository"
	"github.com/ShasiChowdam/user-age-api/internal/routes"
	"github.com/ShasiChowdam/user-age-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	err = loggerpkg.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer loggerpkg.Log.Sync()

	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()

	loggerpkg.Log.Info("Database connected successfully")

	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	routes.SetupRoutes(app, userHandler)

	loggerpkg.Log.Info(
		"Server started",
		zap.String("port", cfg.AppPort),
	)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}