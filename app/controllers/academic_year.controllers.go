package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func AcademicYearList(c *fiber.Ctx) error {
	result := models.AcademicYearList()
	return c.Status(result.Status).JSON(result)
}
