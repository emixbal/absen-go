package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func ClassFetchAll(c *fiber.Ctx) error {
	result := models.ClassFetchAll()
	return c.Status(result.Status).JSON(result)
}
