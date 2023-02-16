package controllers

import (
	"absen-go/app/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FetchAllMembers(c *fiber.Ctx) error {
	limit := 100
	offset := 0
	class_filter := []string{}
	filter_id := ""

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
	if c.Query("filter_id") != "" {
		filter_id = c.Query("filter_id")
	}

	result := models.FethAllMembers(limit, offset, class_filter, filter_id)
	return c.Status(result.Status).JSON(result)
}
