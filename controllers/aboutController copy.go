package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data contact fetched successfully",
		Data:    data.Contact,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}