package handler

import (
	"fmt"

	awssdk "github.com/alsoxavi/go-api-test/aws"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/gofiber/fiber/v2"
)

// SendSMS godoc
// @Summary Sends a new SMS
// @Produce json
// @Tags Messages
// @Accept json
// @Success 201 {object} model.Result{data=model.SMS}
// @Failure 500 {object} model.Result
// @Param sms body model.SMS true "SMS"
// @Security BasicAuth
// @Router /messages [post]
func SendSMS(c *fiber.Ctx) error {
	sms := new(model.SMS)
	if err := c.BodyParser(sms); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Code:    fiber.StatusInternalServerError,
			Error:   !(err != nil),
			Message: err.Error(),
		})
	}

	result, err := awssdk.SendSMS(sms)

	if err != nil || result == "" {
		fmt.Printf("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(model.Result{
			Code:    fiber.StatusInternalServerError,
			Error:   !(err != nil),
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Result{
		Code:    fiber.StatusCreated,
		Error:   !(err != nil),
		Message: result,
	})

}
