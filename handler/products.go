package handler

import (
	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts godoc
// @Summary Returns an array with all products on DB
// @Produce json
// @Tags Products
// @Success 200 {object} model.Result{data=[]model.Product{id=int}}
// @Failure 404 {object} model.Result
// @Security BasicAuth
// @Router /products [get]
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DBConn
	var products []model.Product

	db.Find(&products)

	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(model.Result{
			Code:    fiber.StatusNotFound,
			Error:   true,
			Message: "No products found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Code:  fiber.StatusOK,
		Error: false,
		Data:  products,
	})
}

// GetSingleProduct godoc
// @Summary Returns a single product from DB
// @Produce json
// @Tags Products
// @Success 200 {object} model.Result{data=model.Product{id=int}}
// @Failure 404 {object} model.Result
// @Security BasicAuth
// @Router /products/{id} [get]
// @Param id path integer true "Product ID"
func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product model.Product

	db.Find(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(model.Result{
			Code:    fiber.StatusNotFound,
			Error:   true,
			Message: "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Result{
		Code:  fiber.StatusOK,
		Error: false,
		Data:  product,
	})
}

// CreateProduct godoc
// @Summary Creates a new product in DB
// @Produce json
// @Tags Products
// @Accept json
// @Success 201 {object} model.Result{data=model.Product{id=int}}
// @Failure 500 {object} model.Result
// @Param product body model.Product true "Add product"
// @Security BasicAuth
// @Router /products [post]
func CreateProduct(c *fiber.Ctx) error {
	db := database.DBConn
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Code:    fiber.StatusInternalServerError,
			Error:   !(err != nil),
			Message: err.Error(),
		})
	}

	db.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(model.Result{
		Code:  fiber.StatusCreated,
		Error: false,
		Data:  product,
	})
}

// DeleteProduct godoc
// @Summary Deletes a product from DB
// @Produce json
// @Tags Products
// @Success 204
// @Failure 500 {object} model.Result
// @Security BasicAuth
// @Router /products/{id} [delete]
// @Param id path integer true "Product ID"
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product model.Product

	db.First(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(model.Result{
			Code:    fiber.StatusNotFound,
			Error:   true,
			Message: "Product not found",
		})
	}

	db.Delete(&product)

	return c.Status(fiber.StatusNoContent).JSON(model.Result{
		Code:    fiber.StatusNoContent,
		Error:   false,
		Message: "Product deleted successfully",
	})
}
