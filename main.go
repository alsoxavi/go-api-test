package main

import (
	"log"

	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Fiber Test Rest API
// @version 0.1.0
// @description Starting Swagger docs
// @termsOfService https://version-one.com/terms/

// @contact.name Version-One
// @contact.email support@version.one.com

// @license.name BSD
// @license.url https://opensource.org/licenses/BSD-3-Clause

// @host localhost:3000
// @BasePath /api/v1/

// @securityDefinitions.basic BasicAuth
func main() {
	if err := database.InitDatabase(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)

	log.Println("Starting server")
	app.Listen(":3000")
}
