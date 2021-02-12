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
	v1 := api.Group("/v1")
	products := v1.Group("/products", middleware.AuthReq())
	docs := v1.Group("/docs")
	messages := v1.Group("/messages")

	docs.Get("", swagger.Handler)
	docs.Get("/*", swagger.Handler)
	products.Get("", handler.GetAllProducts)
	products.Get(":id", handler.GetSingleProduct)
	products.Post("", handler.CreateProduct)
	products.Delete(":id", handler.DeleteProduct)
	messages.Post("", handler.SendSMS)
}
