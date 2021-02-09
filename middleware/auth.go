package middleware

import (
	"github.com/alsoxavi/go-api-test/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

// AuthReq works as the auth middleware
func AuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("ADMIN_USERNAME"): config.Config("ADMIN_PASSWORD"),
		},
	}

	err := basicauth.New(cfg)
	return err
}
