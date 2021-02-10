package handler

import (
	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts is a fuction from db
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DBConn
	var products []model.Product

	db.Find(&products)

	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error":   true,
			"message": "No products found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"error":    false,
		"products": products,
	})
}

// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product model.Product

	db.Find(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error":   true,
			"message": "No products found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"error":   false,
		"product": product,
	})
}

// CreateProduct handler
func CreateProduct(c *fiber.Ctx) error {
	db := database.DBConn
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":   true,
			"message": err,
		})
	}

	db.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"error":   false,
		"message": "Product created",
		"product": product,
	})
}

// DeleteProduct from db
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product model.Product

	db.First(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error":   true,
			"message": "No products found",
		})
	}

	db.Delete(&product)

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"error":   false,
		"message": "Product deleted successfully",
	})
}
