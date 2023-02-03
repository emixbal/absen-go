package controllers

import (
	"absen-go/app/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FetchAllStudents(c *fiber.Ctx) error {
	limit := 100
	offset := 0
	class_filter := []string{}

	if c.Query("per_page") != "" {
		limit, _ = strconv.Atoi(c.Query("per_page"))
	}
	if c.Query("page") != "" {
		page, _ := strconv.Atoi(c.Query("page"))
		offset = page - 1
	}
	if c.Query("class") != "" {
		class_filter = strings.Split(c.Query("class"), ",")
	}

	result := models.FethAllStudents(limit, offset, class_filter)
	return c.Status(result.Status).JSON(result)
}
