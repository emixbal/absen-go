package controllers

import (
	"absen-go/app/models"
	"absen-go/app/requests"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func ClassAttendanceMemberArrive(c *fiber.Ctx) error {
	var res models.Response

	p := new(requests.AddClassAttendanceMember)
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

	result, _ := models.ClassAttendanceMemberArrive(p.Code)
	return c.Status(result.Status).JSON(result)
}

func ClassAttendanceMemberLeave(c *fiber.Ctx) error {
	var res models.Response

	p := new(requests.AddClassAttendanceMember)
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

	result, _ := models.ClassAttendanceMemberLeave(p.Code)
	return c.Status(result.Status).JSON(result)
}
