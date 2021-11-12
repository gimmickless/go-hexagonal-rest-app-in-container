package main

import (
	inhttp "github.com/gimmickless/go-hexagonal-rest-app-in-container/internal/transport/ingress/http"
	applog "github.com/gimmickless/go-hexagonal-rest-app-in-containerpkg/logging"
	"github.com/gofiber/fiber/v2"
	httplogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	logger := applog.NewLogger()

	// Load .env file (only expected on local workstations)
	if err := godotenv.Load(".env"); err != nil {
		logger.Debugf(".env file could not be loaded (only harmful when running as standalone on local workstations).")
	}
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(httplogger.New())

	inhttp.Register(app)

	app.Listen(":3000")
}
