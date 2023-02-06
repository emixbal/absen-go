package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func RecapStudentAttendance(c *fiber.Ctx) error {
	class_id := c.Params("class_id")
	month := c.Params("month")

	result := models.RecapStudentAttendance(class_id, month)
	return c.Status(result.Status).JSON(result)
}
