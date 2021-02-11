package handler

import (
	_ "net/http" // Needed for docs

	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts godoc
// @Summary Returns an array with all products on DB
// @Produce json
// @Success 200 {object} model.ProductModel
// @Failure 404 {object} model.ErrorModel
// @Security BasicAuth
// @Router /product [get]
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

// GetSingleProduct godoc
// @Summary Returns a single product from DB
// @Produce json
// @Success 200 {object} model.ProductModel
// @Failure 404 {object} model.ErrorModel
// @Security BasicAuth
// @Router /product/{id} [get]
// @Param id path integer true "Product ID"
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

// CreateProduct godoc
// @Summary Creates a new product in DB
// @Produce json
// @Accept json
// @Success 201 {object} model.ProductModel
// @Failure 500 {object} model.ErrorModel
// @Param product body model.ProductModel true "Add product"
// @Security BasicAuth
// @Router /product [post]
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

// DeleteProduct godoc
// @Summary Deletes a product from DB
// @Produce json
// @Success 204 {object} model.ProductModel
// @Failure 500 {object} model.ErrorModel
// @Security BasicAuth
// @Router /product/{id} [delete]
// @Param id path integer true "Product ID"
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
