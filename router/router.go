package router

import (
	_ "github.com/alsoxavi/go-api-test/docs" // Imported for documentation
	"github.com/alsoxavi/go-api-test/handler"
	"github.com/alsoxavi/go-api-test/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes inits the routs for product access
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", middleware.AuthReq())

	v1.Get("/docs", swagger.Handler)
	v1.Get("/docs/*", swagger.Handler)
	v1.Get("/product", handler.GetAllProducts)
	v1.Get("/product/:id", handler.GetSingleProduct)
	v1.Post("/product", handler.CreateProduct)
	v1.Delete("/product/:id", handler.DeleteProduct)
}
