package controllers

import (
	"absen-go/app/models"

	"github.com/gofiber/fiber/v2"
)

func RecapMemberAttendancePerClass(c *fiber.Ctx) error {
	class_id := c.Params("class_id")
	year_month := c.Params("year_month")

	result := models.RecapMemberAttendancePerClass(class_id, year_month)
	return c.Status(result.Status).JSON(result)
}

func RecapMemberAttendancePerMember(c *fiber.Ctx) error {
	member_id := c.Params("member_id")
	year_month := c.Params("year_month")

	result := models.RecapMemberAttendancePerMember(member_id, year_month)
	return c.Status(result.Status).JSON(result)
}
