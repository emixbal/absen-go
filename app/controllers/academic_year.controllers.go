package controllers

import (
	"absen-go/app/models"
	"absen-go/app/requests"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func AcademicYearList(c *fiber.Ctx) error {
	result := models.AcademicYearList()
	return c.Status(result.Status).JSON(result)
}

func AcademicYearDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	result := models.AcademicYearDetail(id)
	return c.Status(result.Status).JSON(result)
}

func AcademicYearUpdate(c *fiber.Ctx) error {
	var res models.Response

	p := new(requests.AcademicYearUpdateForm)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = "Empty payloads"
		return c.Status(res.Status).JSON(res)
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()
		return c.Status(res.Status).JSON(res)
	}

	id := c.Params("id")

	result := models.AcademicYearUpdate(id, p.Name)
	return c.Status(result.Status).JSON(result)
}
