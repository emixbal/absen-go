package controllers

import (
	"absen-go/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FetchAllStudents(c *fiber.Ctx) error {
	limit := 100
	offset := 0

	if c.Query("per_page") != "" {
		limit, _ = strconv.Atoi(c.Query("per_page"))
	}
	if c.Query("page") != "" {
		offset, _ = strconv.Atoi(c.Query("page"))
		if offset != 0 {
			offset = offset - 1
		}
	}

	result := models.FethAllStudents(limit, offset, c.Query("class"))
	return c.Status(result.Status).JSON(result)
}
