package handler

import (
	"log"

	"github.com/alsoxavi/go-api-test/database"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts is a fuction from db
func GetAllProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result.Products = append(result.Products, product)
	}

	err = c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched product",
		"product": result,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return err
}

// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	defer row.Close()

	row.Next()
	err = row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}

	err = c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Successfully fetched product",
		"product": product,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return err
}

// CreateProduct handler
func CreateProduct(c *fiber.Ctx) error {
	p := new(model.Product)

	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	_, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	err = c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Successfully created product",
		"product": p,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return err
}

// DeleteProduct from db
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}

	log.Println(res)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}
	return err
}
