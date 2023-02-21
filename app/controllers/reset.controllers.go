package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func ResetAttendance(c *fiber.Ctx) error {
	result := models.ResetAttendance()
	return c.Status(result.Status).JSON(result)
}
