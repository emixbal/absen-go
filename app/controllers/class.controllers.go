package controllers

import (
	"absen-go/app/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ClassFetchAll(c *fiber.Ctx) error {
	result := models.ClassFetchAll()
	return c.Status(result.Status).JSON(result)
}

func ClassAddNew(c *fiber.Ctx) error {
	var class models.Class

	payload := struct {
		Name string `json:"name" xml:"name" form:"name" validate:"required"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		log.Println("err BodyParser")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	if len(payload.Name) < 1 {
		log.Println("err payload len < 1")
		return c.Status(400).JSON(fiber.Map{
			"message": "name is required",
		})
	}

	class.Name = payload.Name

	result := models.ClassAddNew(&class)
	return c.Status(result.Status).JSON(result)
}
