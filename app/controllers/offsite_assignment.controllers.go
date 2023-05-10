package controllers

import (
	"absen-go/app/models"
	"absen-go/app/requests"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func OffSiteAssignmentDeparture(c *fiber.Ctx) error {
	var res models.Response

	p := new(requests.OffSiteAssignment)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = "Err payloads"

		return c.Status(res.Status).JSON(res)
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()

		return c.Status(res.Status).JSON(res)
	}

	result, _ := models.OffSiteAssignmentDeparture(p.Code, p.Description)
	return c.Status(result.Status).JSON(result)
}

func OffSiteAssignmentArrive(c *fiber.Ctx) error {
	var res models.Response

	p := struct {
		Code string `json:"code" xml:"code" form:"code" validate:"required"`
	}{}

	if err := c.BodyParser(&p); err != nil {
		log.Println("err BodyParser")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()

		return c.Status(res.Status).JSON(res)
	}

	result, _ := models.OffSiteAssignmentArrive(p.Code)
	return c.Status(result.Status).JSON(result)
}
