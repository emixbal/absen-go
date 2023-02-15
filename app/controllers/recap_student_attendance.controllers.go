package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func RecapMemberAttendance(c *fiber.Ctx) error {
	class_id := c.Params("class_id")
	year_month := c.Params("year_month")

	result := models.RecapMemberAttendance(class_id, year_month)
	return c.Status(result.Status).JSON(result)
}
