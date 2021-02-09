package router

import (
	"github.com/alsoxavi/go-api-test/handler"
	"github.com/alsoxavi/go-api-test/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes inits the routs for product access
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", middleware.AuthReq())

	v1.Get("/product", handler.GetAllProducts)
	v1.Get("/product/:id", handler.GetSingleProduct)
	v1.Post("/product", handler.CreateProduct)
	v1.Delete("/product/:id", handler.DeleteProduct)
}
