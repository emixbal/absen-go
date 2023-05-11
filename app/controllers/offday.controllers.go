package controllers

import (
	"absen-go/app/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func OffdayNew(c *fiber.Ctx) error {
	var res models.Response

	p := struct {
		Date string `json:"date" xml:"date" form:"date" validate:"required"`
	}{}

	if err := c.BodyParser(&p); err != nil {
		log.Println("err BodyParser")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	t, err := time.Parse("01/02/2006", p.Date)
	if err != nil {
		log.Println("OffdayNew controller parse date error")
		log.Println(err)

		res.Status = 400
		res.Message = "parse date error, check your date format. Accepted: MM/DD/YYYY"
		return c.Status(400).JSON(res)
	}

	result := models.OffdayAddNew(t)
	return c.Status(result.Status).JSON(result)
}

func OffdatFechAll(c *fiber.Ctx) error {
	result, _ := models.OffdayFetchAll()
	return c.Status(result.Status).JSON(result)
}

func OffdayHardDelete(c *fiber.Ctx) error {
	result, _ := models.OffdayHardDelete(c.Params("id"))

	return c.Status(result.Status).JSON(result)
}