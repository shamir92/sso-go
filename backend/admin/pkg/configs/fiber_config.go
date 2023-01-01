package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	appName := os.Getenv("APP_NAME")

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     appName,
		Prefork:     true,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
