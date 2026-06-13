package middleware

import (
	"time"

	"github.com/ShasiChowdam/user-age-api/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info(
	"Request completed",
	zap.String("method", c.Method()),
	zap.String("path", c.Path()),
	zap.Duration("duration", duration),
)

		return err
	}
}