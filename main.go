package main

import (
	"log"

	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)

	log.Println("Starting server")
	app.Listen(":3000")
}
