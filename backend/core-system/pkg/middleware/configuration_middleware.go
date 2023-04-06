package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func ConfigurationProtected() func(ctx *fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	// if ctx.request

	// return jwtMiddleware.New(config)
	return nil
}
