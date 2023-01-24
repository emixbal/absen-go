package controllers

import (
	"absen-go/app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddExtracurricularAttendanceStudent(c *fiber.Ctx) error {
	var book models.Book

	book.Name = c.FormValue("name")

	if book.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}

	result, _ := models.AddExtracurricularAttendanceStudent(&book)
	return c.Status(result.Status).JSON(result)
}
